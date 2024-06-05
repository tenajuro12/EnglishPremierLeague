package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"news_service/internal/entity"
	"sort"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	pb "news_service/internal/proto"

	"news_service/internal/interfaces/repository"
	"news_service/internal/usecases"
)

type server struct {
	pb.UnimplementedNewsServiceServer
	articles      []*pb.Article
	createUseCase *usecases.CreateNewsArticle
	getUseCase    *usecases.GetNewsArticle
	updateUseCase *usecases.UpdateNewsArticle
	deleteUseCase *usecases.DeleteNewsArticle
	listUseCase   *usecases.ListNewsArticles
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123456@localhost/news?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewNewsArticlePostgresRepository(db)

	createUseCase := &usecases.CreateNewsArticle{Repository: repo}
	getUseCase := &usecases.GetNewsArticle{Repository: repo}
	updateUseCase := &usecases.UpdateNewsArticle{Repository: repo}
	deleteUseCase := &usecases.DeleteNewsArticle{Repository: repo}
	listUseCase := &usecases.ListNewsArticles{Repository: repo}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNewsServiceServer(grpcServer, &server{
		createUseCase: createUseCase,
		getUseCase:    getUseCase,
		updateUseCase: updateUseCase,
		deleteUseCase: deleteUseCase,
		listUseCase:   listUseCase,
	})
	log.Println("Starting gRPC server on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	article := entity.NewsArticle{
		Title:     req.GetTitle(),
		Content:   req.GetContent(),
		Category:  req.GetCategory(),
		Author:    req.GetAuthor(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdArticle, err := s.createUseCase.Execute(article)
	if err != nil {
		return nil, err
	}

	return &pb.CreateArticleResponse{Id: int32(createdArticle.ID)}, nil
}

func (s *server) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	article, err := s.getUseCase.Execute(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.GetArticleResponse{
		Id:        int32(article.ID),
		Title:     article.Title,
		Content:   article.Content,
		Category:  article.Category,
		Author:    article.Author,
		CreatedAt: article.CreatedAt.String(),
		UpdatedAt: article.UpdatedAt.String(),
	}, nil
}

func (s *server) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	article := entity.NewsArticle{
		ID:       int(req.GetId()),
		Title:    req.GetTitle(),
		Content:  req.GetContent(),
		Category: req.GetCategory(),
		Author:   req.GetAuthor(),
	}

	updatedArticle, err := s.updateUseCase.Execute(article)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateArticleResponse{
		Id:        int32(updatedArticle.ID),
		Title:     updatedArticle.Title,
		Content:   updatedArticle.Content,
		Category:  updatedArticle.Category,
		Author:    updatedArticle.Author,
		CreatedAt: updatedArticle.CreatedAt.String(),
		UpdatedAt: updatedArticle.UpdatedAt.String(),
	}, nil
}

func (s *server) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	err := s.deleteUseCase.Execute(int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteArticleResponse{}, nil
}

func (s *server) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
	articles, err := s.listUseCase.Execute(ctx, &usecases.ListNewsArticles{})
	if err != nil {
		return nil, err
	}

	filteredArticles := articles

	if req.GetFilterByAuthor() != "" {
		filteredArticles = filterArticlesByAuthor(filteredArticles, req.GetFilterByAuthor())
	}

	sortedArticles := filteredArticles
	if req.GetSortBy() != "" {
		sortedArticles = sortArticles(filteredArticles, req.GetSortBy(), req.GetSortOrder())
	}

	pagedArticles := paginateArticles(sortedArticles, req.GetPage(), req.GetPageSize())

	var pbArticles []*pb.Article
	for _, article := range pagedArticles {
		pbArticle := &pb.Article{
			Id:        int32(article.ID),
			Title:     article.Title,
			Content:   article.Content,
			Category:  article.Category,
			Author:    article.Author,
			CreatedAt: article.CreatedAt.String(),
			UpdatedAt: article.UpdatedAt.String(),
		}
		pbArticles = append(pbArticles, pbArticle)
	}

	return &pb.ListArticlesResponse{
		Articles:   pbArticles,
		TotalCount: int32(len(sortedArticles)),
	}, nil
}

func filterArticlesByAuthor(articles []*entity.NewsArticle, author string) []*entity.NewsArticle {
	if author == "" {
		return articles
	}
	var filteredArticles []*entity.NewsArticle
	for _, article := range articles {
		if article.Author == author {
			filteredArticles = append(filteredArticles, article)
		}
	}
	return filteredArticles
}

func sortArticles(articles []*entity.NewsArticle, sortBy, sortOrder string) []*entity.NewsArticle {
	switch sortBy {
	case "title":
		sort.Slice(articles, func(i, j int) bool {
			if sortOrder == "desc" {
				return articles[i].Title > articles[j].Title
			}
			return articles[i].Title < articles[j].Title
		})
	case "category":
		sort.Slice(articles, func(i, j int) bool {
			if sortOrder == "desc" {
				return articles[i].Category > articles[j].Category
			}
			return articles[i].Category < articles[j].Category
		})
	case "author":
		sort.Slice(articles, func(i, j int) bool {
			if sortOrder == "desc" {
				return articles[i].Author > articles[j].Author
			}
			return articles[i].Author < articles[j].Author
		})
	case "created_at":
		sort.Slice(articles, func(i, j int) bool {
			if sortOrder == "desc" {
				return articles[i].CreatedAt.After(articles[j].CreatedAt)
			}
			return articles[i].CreatedAt.Before(articles[j].CreatedAt)
		})
	case "updated_at":
		sort.Slice(articles, func(i, j int) bool {
			if sortOrder == "desc" {
				return articles[i].UpdatedAt.After(articles[j].UpdatedAt)
			}
			return articles[i].UpdatedAt.Before(articles[j].UpdatedAt)
		})
	}
	return articles
}

func paginateArticles(articles []*entity.NewsArticle, page, pageSize int32) []*entity.NewsArticle {
	start := int((page - 1) * pageSize)
	end := int(page * pageSize)
	if start >= len(articles) {
		return []*entity.NewsArticle{}
	}
	if end > len(articles) {
		end = len(articles)
	}
	return articles[start:end]
}
