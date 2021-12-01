package com.example.zoo.service;

import com.example.zoo.domain.Animals;
import com.example.zoo.repos.AnimalsRepo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Locale;

@Service
public class AnimalService {
    @Autowired
    private final AnimalsRepo animalsRepo;

    public AnimalService(AnimalsRepo addressRepository){
        this.animalsRepo = addressRepository;
    }

    public void createAnimals(Animals animals){

        animalsRepo.save(animals);
    }

    public void deleteAnimal(Animals animals){
        animalsRepo.delete(animals);
    }
    public void deleteAnimals(){
        animalsRepo.deleteAll();
    }

    public List<Animals> findAll(){
        return animalsRepo.findAll();
    }

    public Animals findById(Long userId){
        return animalsRepo.findById(userId).orElse(null);
    }

    public List<Animals> findAllByName(String name){
        return animalsRepo.findAllByName(name.toLowerCase(Locale.ROOT));
    }
}
