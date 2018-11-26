package handlers

import (
	"context"
	"github.com/jinzhu/gorm"
	pb "yekGoProtoBuffer/server/generated"
)

type OgrenciHandler struct {
	db *gorm.DB
}

func NewOgrenciHandler(db *gorm.DB) OgrenciHandler {
	return OgrenciHandler{db}
}

func (h *OgrenciHandler) DeleteOgrenciByID(c context.Context, input *pb.OgrenciID) (output *pb.OgrenciReply, err error) {
	var ogrenci pb.Ogrenci
	if err := h.db.Where("id = ?", &input.Id).First(&ogrenci).Error; err != nil {
		return nil, err
	}

	h.db.Delete(&ogrenci)
	return &pb.OgrenciReply{Ogrenci: &ogrenci}, nil
}

func (h *OgrenciHandler) UpdateOgrenci(c context.Context, input *pb.OgrenciRequest) (output *pb.OgrenciReply, err error) {
	var ogrenci pb.Ogrenci
	if err := h.db.Where("id = ?", &input.Ogrenci.Id).First(&ogrenci).Error; err != nil {
		return nil, err
	}

	if err := h.db.Save(&input.Ogrenci).Error; err != nil {
		return nil, err
	}

	return &pb.OgrenciReply{Ogrenci: &ogrenci}, nil
}

func (h *OgrenciHandler) InsertOgrenci(c context.Context, input *pb.OgrenciRequest) (output *pb.OgrenciReply, err error) {

	var ogrenci *pb.Ogrenci
	if err := h.db.Create(&input.Ogrenci).Error; err != nil {
		return nil, err
	}

	return &pb.OgrenciReply{Ogrenci: ogrenci}, nil
}

func (h *OgrenciHandler) GetOgrenciByID(c context.Context, input *pb.OgrenciID) (output *pb.OgrenciReply, err error) {

	var ogrenci pb.Ogrenci
	if err := h.db.Where("id = ?", &input.Id).First(&ogrenci).Error; err != nil {
		return nil, err
	}

	return &pb.OgrenciReply{Ogrenci: &ogrenci}, nil
}

func (h *OgrenciHandler) GetOgrenciAll(c context.Context, input *pb.OgrenciID) (output *pb.OgrenciListReply, err error) {
	var ogrenci []*pb.Ogrenci
	if err := h.db.Find(&ogrenci).Error; err != nil {
		return nil, err
	}

	return &pb.OgrenciListReply{Ogrenci: ogrenci}, nil
}
