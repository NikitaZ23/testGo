package com.example.zoo.service;

import com.example.zoo.domain.Products;
import com.example.zoo.repos.ProductsRepo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Locale;

@Service
public class ProductsService{
    @Autowired
    private final ProductsRepo productsRepo;

    public ProductsService(ProductsRepo productsRepo){
        this.productsRepo = productsRepo;
    }

    public void createProduct(Products products){
        productsRepo.save(products);
    }

    public void deleteProduct(Products products){
        productsRepo.delete(products);
    }

    public void deleteAll(){
        productsRepo.deleteAll();
    }

    public List<Products> findAll(){
        return productsRepo.findAll();
    }

    public Products findById(Long userId){
        return productsRepo.findById(userId).orElse(null);
    }

    public List<Products> findAllByName(String name){
        return productsRepo.findAllByName(name.toLowerCase(Locale.ROOT));
    }
}
