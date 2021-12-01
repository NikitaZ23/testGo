package com.example.zoo.repos;

import com.example.zoo.domain.Animals;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface AnimalsRepo extends JpaRepository<Animals,Long> {
    List<Animals> findAllByName(String name);//просто правильное название метода даст возможность
    //избежать запросов на SQL

//    @Query(value = "select * from Animals where name like '%о%'", nativeQuery = true)
//        //если и этого мало - можно написать запрос на чистом SQL и все это будет работать
//    List<Animals> findWhereNameStartsFromSmith();
}
