package com.example.zoo.service;

import com.example.zoo.domain.SOP;
import com.example.zoo.repos.SopRepo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class SopService {
    @Autowired
    private final SopRepo sopRepo;

    public SopService(SopRepo sopRepo){
        this.sopRepo = sopRepo;
    }

    public void createSOP(SOP sop){
        if (findByIds(sop.getIda(), sop.getIdp()) == null)
            sopRepo.save(sop);
    }

    public void updateSOP(SOP sop){
            sopRepo.save(sop);
    }

    public void deleteSOP(SOP sop){
        sopRepo.delete(sop);
    }

    public void deleteAll(){
        sopRepo.deleteAll();
    }

    public List<SOP> findAll(){
        return sopRepo.findAll();
    }

    public List<SOP> findByIda(int ida){
        return sopRepo.findByIda(ida);
    }

    public List<SOP> findByIdp(int idp){
        return sopRepo.findByIdp(idp);
    }

    public SOP findByIds(Long ida, Long idp){
        return sopRepo.findByIds(ida, idp);
    }

    public SOP findById(Long userId){
        return sopRepo.findById(userId).orElse(null);
    }

    public List<SOP> findByNameP(String name){
        System.out.println(name);
        return sopRepo.findAllByNameP(name);
    }
}
