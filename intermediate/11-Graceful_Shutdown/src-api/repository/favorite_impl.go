package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-class/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RealFavoriteRepository struct {
	db *pgxpool.Pool
}

func (r *RealFavoriteRepository) GetFavoriteByID(ctx context.Context, id string) (*model.Favorite, error) {
	var favorite model.Favorite
	err := r.db.QueryRow(
		ctx,
		"SELECT id, image_url, created_at FROM favorites WHERE id = $1",
		id,
	).Scan(&favorite.ID, &favorite.ImageUrl, &favorite.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("favorite not found")
		}
		return nil, fmt.Errorf("query failed: %v", err)
	}

	return &favorite, nil
}

func (r *RealFavoriteRepository) DeleteFavoriteByID(ctx context.Context, id string) (*model.Favorite, error) {
	var favorite model.Favorite
	err := r.db.QueryRow(
		ctx,
		"DELETE FROM favorites WHERE id = $1 RETURNING id, image_url, created_at",
		id,
	).Scan(&favorite.ID, &favorite.ImageUrl, &favorite.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("favorite not found")
		}
		return nil, fmt.Errorf("delete failed: %v", err)
	}

	return &favorite, nil
}

func (r *RealFavoriteRepository) InsertFavorite(ctx context.Context, imageUrl string) (*model.Favorite, error) {
	var favorite model.Favorite
	err := r.db.QueryRow(
		ctx,
		"INSERT INTO favorites (image_url) VALUES ($1) RETURNING id, image_url, created_at",
		imageUrl,
	).Scan(&favorite.ID, &favorite.ImageUrl, &favorite.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("insert failed: %v", err)
	}
	return &favorite, nil
}

func (r *RealFavoriteRepository) GetAllFavorites(ctx context.Context) ([]model.Favorite, error) {
	rows, err := r.db.Query(ctx, "SELECT id, image_url, created_at FROM favorites")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var favorites []model.Favorite
	for rows.Next() {
		var fav model.Favorite
		err := rows.Scan(&fav.ID, &fav.ImageUrl, &fav.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %v", err)
		}
		favorites = append(favorites, fav)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return favorites, nil
}

func NewRealFavoriteRepository(pool *pgxpool.Pool) FavoriteRepository {
	return &RealFavoriteRepository{
		db: pool,
	}
}
