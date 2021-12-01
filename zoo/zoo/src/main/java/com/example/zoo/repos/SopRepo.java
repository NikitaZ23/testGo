package com.example.zoo.repos;

import com.example.zoo.domain.SOP;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;

public interface SopRepo extends JpaRepository<SOP, Long> {

    @Query(value = "select * from sop where ida = ?1", nativeQuery = true)
    List<SOP> findByIda(int ida);

    @Query(value = "select * from sop where idp = ?1", nativeQuery = true)
    List<SOP> findByIdp(int idp);

    @Query(value = "select * from sop where ida = ?1 and idp = ?2", nativeQuery = true)
    SOP findByIds(Long ida, Long idp);

    @Query(value = "select s.* from sop s, products p where s.idp = p.id and p.name = ?1", nativeQuery = true)
    List<SOP> findAllByNameP(@Param("name") String name);
}
