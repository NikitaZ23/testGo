package com.example.zoo.Controls;

import com.example.zoo.domain.Animals;
import com.example.zoo.domain.Products;
import com.example.zoo.domain.SOP;
import com.example.zoo.service.AnimalService;
import com.example.zoo.service.ProductsService;
import com.example.zoo.service.SopService;
import com.fasterxml.jackson.core.JsonProcessingException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;


import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;


@RestController
@RequestMapping("/api")
public class MainControl {

    @Autowired
    private AnimalService animalService;

    @Autowired
    private ProductsService productsService;

    @Autowired
    private SopService sopService;

    @GetMapping("/animals")
    List<Animals> getAllAnimals() {
        return animalService.findAll();
    }

    @GetMapping("/animal")
    Animals getAnimal(@RequestParam(value = "id", required = true) Long id) {
        return animalService.findById(id);
    }

    @PostMapping("/animal")
    ResponseEntity<Void> createAnimal(@RequestBody Animals animals) {
        System.out.println(animals);
        animalService.createAnimals(animals);
        return ResponseEntity.ok().build();
    }

    @DeleteMapping("/animal")
    ResponseEntity<Void> delAnimal(@RequestParam(value = "id", required = true) Long id) {
        for(SOP s: sopService.findByIda(Math.toIntExact(id)))
            sopService.deleteSOP(s);
        animalService.deleteAnimal(animalService.findById(id));
        return ResponseEntity.ok().build();
    }

    @DeleteMapping("/animals")
    ResponseEntity<Void> delAnimals() {
        System.out.println("delAnimals");
        animalService.deleteAnimals();
        sopService.deleteAll();
        return ResponseEntity.ok().build();
    }

    @DeleteMapping("/animalSP")
    ResponseEntity<Void> delAnimals2(@RequestParam("ids") List<String> ids) {
         for(String id: ids) {
             for (SOP s : sopService.findByIda(Integer.valueOf(id)))
                sopService.deleteSOP(s);
            animalService.deleteAnimal(animalService.findById(Long.valueOf(id)));
        }
        return ResponseEntity.ok().build();
    }

///////////////////////////////////////////////////////////////////////////////////////////////////////////
    @GetMapping("/products")
    List<Products> getAllProducts() {
        return productsService.findAll();
    }

    @GetMapping("/product")
    Products getProduct(@RequestParam(value = "id", required = true) Long id) {
        return productsService.findById(id);
    }

    @PostMapping("/product")
    ResponseEntity<Void> createProduct(@RequestBody Products products) {
        System.out.println(products);
        productsService.createProduct(products);
        return ResponseEntity.ok().build();
    }

    @DeleteMapping("/product")
    ResponseEntity<Void> delProduct(@RequestParam(value = "id", required = true) Long id) {
        productsService.deleteProduct(productsService.findById(id));

        for(SOP s: sopService.findByIdp(Math.toIntExact(id)))
            sopService.deleteSOP(s);

        return ResponseEntity.ok().build();
    }

    @DeleteMapping("/products")
    ResponseEntity<Void> delProducts() {
        productsService.deleteAll();
        sopService.deleteAll();
        return ResponseEntity.ok().build();
    }

    @DeleteMapping("/productSP")
    ResponseEntity<Void> delProducts2(@RequestParam("ids") List<String> ids) {
        for(String id: ids) {
            for(SOP s: sopService.findByIdp(Integer.valueOf(id)))
                sopService.deleteSOP(s);
            productsService.deleteProduct(productsService.findById(Long.valueOf(id)));
        }
        return ResponseEntity.ok().build();
    }

    @PutMapping("/product")
    ResponseEntity<Void> editCurProduct(@RequestParam(value = "id", required = true) Long id, @RequestParam(value = "cur", required = true) int cur) {
        Products products = productsService.findById(id);
        products.setCurnumber(cur);
        productsService.createProduct(products);
        return ResponseEntity.ok().build();
    }
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

    @GetMapping("/7days")
    ResponseEntity<List<HashMap<String, Object>>> get7Days(@RequestParam(value = "data", required = true) String data) throws JsonProcessingException {

        List<Products> pr =  productsService.findAll();

        List<HashMap<String, Object>> map = new ArrayList<>();
        for(Products p: pr){
            List<SOP> s = sopService.findByNameP(p.getName());
            double sum =  0;
            for(SOP sl: s)
                sum += sl.getNumber();

            sum*=7;
            System.out.println("sum:" + sum + " storage: " + p.getCurnumber());

            HashMap<String, Object> m = new HashMap<>();

            m.put("Название", p.getName());
            m.put("Сколько требуется", sum);
            m.put("Сколько есть", p.getCurnumber());

            if(getRaz(p.getCurnumber(), sum) > 0)
                m.put("Не хватает", getRaz(p.getCurnumber(), sum));
            else
                m.put("Не хватает", null);
            m.put("Ед. измерения", p.getUnit());

            map.add(m);
        }

        return ResponseEntity.ok(map);
    }

    double getRaz(double stor, double sum){
        double raz;

        raz = sum - stor;
        System.out.println(raz);
        return raz;
    }

    @GetMapping("/animalsFull")
    ResponseEntity<List<HashMap<String, Object>>> anFull(){
        List<HashMap<String, Object>> list = new ArrayList<>();

        List<Animals> animals = animalService.findAll();

        for(Animals a: animals){
            HashMap<String, Object> map = new HashMap<>();
            map.put("Наименование", a.getName());
            map.put("Вид", a.getView());
            System.out.println(sopService.findByIda(Math.toIntExact(a.getId())).size());
            for(SOP s: sopService.findByIda(Math.toIntExact(a.getId()))) {
                HashMap<String, Object> map2 = new HashMap<>();
                map2.put("Имя продукта", productsService.findById(s.getIdp()).getName());
                map2.put("Тип продукта", productsService.findById(s.getIdp()).getType());
                map2.put("Потребляет", s.getNumber());
                map2.put("Ед. измерения", productsService.findById(s.getIdp()).getUnit());
                map.put("Продукт id: " + productsService.findById(s.getIdp()).getId(),map2);
            }

            list.add(map);
        }

        return ResponseEntity.ok(list);
    }
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    @PostMapping("/sop")
    ResponseEntity<Void> createSOP(@RequestBody SOP sop) {
        sopService.createSOP(sop);
        return ResponseEntity.ok().build();
    }

    @PutMapping("/sop")
    ResponseEntity<Void> editAnimalSOP(@RequestParam(value = "ida", required = true) Long ida,
                                       @RequestParam(value = "idp", required = true) Long idp,
                                       @RequestParam(value = "number", required = true) double cur) {

        SOP sop = sopService.findByIds(ida, idp);
        sop.setNumber(cur);
        sopService.updateSOP(sop);
        return ResponseEntity.ok().build();
    }
}
