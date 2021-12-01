package com.example.zoo;

import com.example.zoo.domain.Animals;
import com.example.zoo.domain.Products;
import com.example.zoo.domain.SOP;
import com.example.zoo.service.AnimalService;
import com.example.zoo.service.ProductsService;
import com.example.zoo.service.SopService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.event.EventListener;
import org.springframework.data.util.Pair;

import java.sql.*;
import java.util.*;



@SpringBootApplication
public class ZooApplication {

	@Autowired
	private AnimalService animalService;

	@Autowired
	private ProductsService productService;

	@Autowired
	private SopService sopService;

	public static void main(String[] args) throws SQLException {
		SpringApplication.run(ZooApplication.class, args);
	}

	@EventListener(ApplicationReadyEvent .class)
	private void testJpaMethods() throws SQLException {
		createDB();

		ArrayList<Animals> list = new ArrayList<>();
		list.add(new Animals("Обезьяна", "млекопитающие", false));
		list.add(new Animals("Заяц", "млекопитающие", false));
		list.add(new Animals("Орел", "птица", true));
		list.add(new Animals("Тигр", "млекопитающие", true));
		list.add(new Animals("Лошадь", "млекопитающие", false));

		for(Animals str: list){
			List<Animals> l = animalService.findAllByName(str.getName());
			if(l.size() == 0)
				animalService.createAnimals(str);
		}

		animalService.findAll().forEach(it-> System.out.println(it.getName()));

		animalService.findAllByName("Обезьяна").forEach(it-> System.out.println(it.getName()));
		System.out.println("------------------------------------------------------------------");

		ArrayList<Products> list2 = new ArrayList<>();
		list2.add(new Products("Морковь", 100, "шт", "овощ"));
		list2.add(new Products("Мясо", 10, "кг", "мясо"));
		list2.add(new Products("Капуста", 40, "шт", "овощ"));
		list2.add(new Products("Банан", 50, "шт", "фрукт"));
		list2.add(new Products("Яблоко", 0, "шт", "фрукт"));
		list2.add(new Products("Зерно", 1000, "кг", "овощь"));
		list2.add(new Products("Вода", 100, "л", "жидкость"));

		for(Products pr: list2){
			List<Products> p = productService.findAllByName(pr.getName());
			if(p.size() == 0)
				productService.createProduct(pr);
		}

		productService.findAll().forEach(it-> System.out.println(it.getName()));

		System.out.println("------------------------------------------------------------------");

		List<SOP> s = new ArrayList<>();
		s.add(new SOP(animalService.findAllByName("Обезьяна").get(0).getId(), productService.findAllByName("Вода").get(0).getId(),1.0));
		s.add(new SOP(animalService.findAllByName("Обезьяна").get(0).getId(), productService.findAllByName("банан").get(0).getId(),10.0));
		s.add(new SOP(animalService.findAllByName("Обезьяна").get(0).getId(), productService.findAllByName("яблоко").get(0).getId(),5.0));


		s.add(new SOP(animalService.findAllByName("Заяц").get(0).getId(), productService.findAllByName("Вода").get(0).getId(),1.0));
		s.add(new SOP(animalService.findAllByName("Заяц").get(0).getId(), productService.findAllByName("Морковь").get(0).getId(),10.0));
		s.add(new SOP(animalService.findAllByName("Заяц").get(0).getId(), productService.findAllByName("Капуста").get(0).getId(),5.0));

		s.add(new SOP(animalService.findAllByName("Орел").get(0).getId(), productService.findAllByName("Вода").get(0).getId(),1.0));
		s.add(new SOP(animalService.findAllByName("Орел").get(0).getId(), productService.findAllByName("Мясо").get(0).getId(),10.0));
		s.add(new SOP(animalService.findAllByName("Орел").get(0).getId(), productService.findAllByName("Зерно").get(0).getId(),5.0));
		s.add(new SOP(animalService.findAllByName("Орел").get(0).getId(), productService.findAllByName("яблоко").get(0).getId(),5.0));

		s.add(new SOP(animalService.findAllByName("Тигр").get(0).getId(), productService.findAllByName("Вода").get(0).getId(),1.0));
		s.add(new SOP(animalService.findAllByName("Тигр").get(0).getId(), productService.findAllByName("Мясо").get(0).getId(),10.0));

		s.add(new SOP(animalService.findAllByName("Лошадь").get(0).getId(), productService.findAllByName("Вода").get(0).getId(),1.0));
		s.add(new SOP(animalService.findAllByName("Лошадь").get(0).getId(), productService.findAllByName("Зерно").get(0).getId(),10.0));
		s.add(new SOP(animalService.findAllByName("Лошадь").get(0).getId(), productService.findAllByName("яблоко").get(0).getId(),5.0));
		s.add(new SOP(animalService.findAllByName("Лошадь").get(0).getId(), productService.findAllByName("Морковь").get(0).getId(),5.0));

		for(SOP sf: s)
			sopService.createSOP(sf);

		sopService.findAll().forEach(it-> System.out.println(animalService.findById(it.getIda()).getName() + " | " + productService.findById(it.getIdp()).getName()+ " | " + it.getNumber().toString()));
	}

	private static void createDB() throws SQLException {
		Connection c = DriverManager.getConnection("jdbc:postgresql://localhost:5432/zoo", "root", "1234");
		if (c!= null){
			System.out.println("Connection created successfully");

		}else{
			System.out.println("Problem with creating connection");
		}
		Statement st = c.createStatement();
		int res;

		try {
			ResultSet result = st.executeQuery("SELECT * FROM animals");
		}
		catch (Exception e) {
			try {
				res = st.executeUpdate("CREATE TABLE animals (" +
						"id INT NOT NULL," +
						"Name VARCHAR(45) NOT NULL," +
						"view VARCHAR(45) NOT NULL," +
						"predator BOOLEAN NOT NULL," +
						"PRIMARY KEY (id));");
			} catch (Exception e2) {
				e2.printStackTrace();
			}
		}

		try {
			ResultSet result = st.executeQuery("SELECT * FROM products");
		}
		catch (Exception e) {
			try {
				res = st.executeUpdate("CREATE TABLE products (" +
						"  id INT NOT NULL," +
						"  Name VARCHAR(45) NOT NULL," +
						"  CurNumber INT NOT NULL," +
						"  Unit VARCHAR(45) NOT NULL," +
						"  Type VARCHAR(45) NOT NULL," +
						"PRIMARY KEY (id));");
			} catch (Exception e2) {
				e2.printStackTrace();
			}
		}

		try {
			ResultSet result = st.executeQuery("SELECT * FROM sop");
		}
		catch (Exception e) {
			try {
				res = st.executeUpdate("CREATE TABLE sop (" +
						"id INT NOT NULL," +
						"idA INT NOT NULL," +
						"idP INT NOT NULL," +
						"number double NOT NULL," +
						"PRIMARY KEY (id), "+
						"CONSTRAINT idAn FOREIGN KEY (idA)" +
						"REFERENCES animals (id)," +
						"CONSTRAINT idPr " +
						"FOREIGN KEY (idP)" +
						"REFERENCES products (id));");
			} catch (Exception e2) {
				e2.printStackTrace();
			}
		}
	}
}
