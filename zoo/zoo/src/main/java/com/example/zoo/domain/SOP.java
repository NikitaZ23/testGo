package com.example.zoo.domain;

import javax.persistence.*;
import java.util.Objects;

@Entity
@Table(name = "sop")
public class SOP {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private int id;

    @Column(name = "ida")
    Long ida;

    @Column(name = "idp")
    Long idp;

    @Column(name = "number")
    Double number;

    public SOP() {
    }

    public SOP(Long ida, Long idp, Double number) {
        this.ida = ida;
        this.idp = idp;
        this.number = number;
    }

    public int getId() {
        return id;
    }

    public Double getNumber() {
        return number;
    }

    public void setNumber(Double number) {
        this.number = number;
    }

    public Long getIda() {
        return ida;
    }

    public void setIda(Long ida) {
        this.ida = ida;
    }

    public Long getIdp() {
        return idp;
    }

    public void setIdp(Long idp) {
        this.idp = idp;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        SOP sop = (SOP) o;
        return (ida==sop.ida) && (idp==sop.idp);
    }

    @Override
    public int hashCode() {
        return Objects.hash(ida, idp);
    }
}
