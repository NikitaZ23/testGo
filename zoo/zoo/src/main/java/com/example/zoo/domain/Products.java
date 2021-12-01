package com.example.zoo.domain;

import javax.persistence.*;
import java.util.Locale;

@Entity
@Table(name = "Products")
public class Products {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    @Column(name = "name")
    String name;

    @Column(name = "curnumber")
    double curnumber;

    @Column(name = "unit")
    String unit;

    @Column(name = "type")
    String type;

    public Products() {
    }

    public Products(String name, double curnumber, String unit, String type) {
        this.name = name.toLowerCase(Locale.ROOT);
        this.curnumber = curnumber;
        this.unit = unit.toLowerCase(Locale.ROOT);
        this.type = type.toLowerCase(Locale.ROOT);
    }

    public Long getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public double getCurnumber() {
        return curnumber;
    }

    public void setCurnumber(int curnumber) {
        this.curnumber = curnumber;
    }

    public String getUnit() {
        return unit;
    }

    public void setUnit(String unit) {
        this.unit = unit;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }
}
