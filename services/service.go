package services

import (
	"go-postgres-menu/helper"
	"go-postgres-menu/models"
	"go-postgres-menu/repository"
)

func Create(request models.Request) models.Master {
	master := models.Master{
		Name: request.Name,
	}
	lastIdMaster := repository.Insert(master)

	categoryReq := models.Category{
		MasterId:     lastIdMaster,
		NameCategory: request.NameCategory,
	}
	lastIdCategory := repository.InsertCategory(categoryReq)

	for _, nameBahan := range request.NameBahan {
		bahanReq := models.Bahan{
			MasterId:   lastIdMaster,
			CategoryId: lastIdCategory,
			NameBahan:  nameBahan,
		}
		repository.InsertBahan(bahanReq)
	}

	return master
}

func GetAll(name string) []models.ResponseResult {
	menus, err := repository.GetAllMenu(name)
	if err != nil {
		panic(err)
	}
	return menus
}

func GetMenu(id int64) (models.ResponseResult, error) {
	master, err := repository.GetMenu(id)
	if err != nil {
		panic(err)
	}
	return master, err
}

func UpdateMenu(id int64, request models.Request) (models.ResponseResult, error) {
	menu, err := repository.GetMenu(id)
	if err != nil {
		helper.Panic(err)
	}

	master := models.Master{
		Name: request.Name,
	}
	masterID := repository.UpdateMenu(id, master)

	categoryReq := models.Category{
		NameCategory: request.NameCategory,
	}
	categoryID := repository.UpdateCategory(id, categoryReq)

	bahanReq := request.NameBahan
	if len(bahanReq) > 0 {
		err := repository.DeleteBahanByMasterID(id)
		if err != nil {
			helper.Panic(err)
		}

		for _, bahanName := range bahanReq {
			bahan := models.Bahan{
				NameBahan:  bahanName,
				CategoryId: categoryID,
				MasterId:   masterID,
			}
			err := repository.CreateBahan(bahan)
			if err != nil {
				return models.ResponseResult{}, err
			}
		}
	}

	return menu, err
}

func DeleteMenu(id int64) error {
	err := repository.DeleteBahan(id)
	if err != nil {
		helper.Panic(err)
	}

	err = repository.DeleteCategory(id)
	if err != nil {
		helper.Panic(err)
	}

	err = repository.DeleteMenu(id)
	if err != nil {
		helper.Panic(err)
	}

	return err
}
