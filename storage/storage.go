package storage

import (
	"github.com/benrobinson-ms/go-blog-api/database"
	"github.com/benrobinson-ms/go-blog-api/models"
)

func GetAllArticles() ([]models.Article, error) {
	var articles []models.Article
	if err := database.DB.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func GetArticleByID(id uint) (*models.Article, error) {
	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func CreateArticle(article *models.Article) error {
	if err := database.DB.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func UpdateArticle(article *models.Article) error {
	if err := database.DB.Save(article).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticleByID(id uint) error {
	if err := database.DB.Delete(&models.Article{}, id).Error; err != nil {
		return err
	}
	return nil
}
