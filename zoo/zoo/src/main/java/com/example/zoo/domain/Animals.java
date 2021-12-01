package com.example.zoo.domain;

import javax.persistence.*;
import java.util.Locale;

@Entity
@Table(name = "animals")
public class Animals {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
     private Long id;

    @Column(name = "name")
    String name;

    @Column(name = "view")
    String view;

    @Column(name = "predator")
    Boolean predator;

    public Animals() {
    }

    public Animals(String name, String view, Boolean predator) {
        this.name = name.toLowerCase(Locale.ROOT);
        this.view = view.toLowerCase(Locale.ROOT);
        this.predator = predator;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Long getId() {
        return id;
    }

    public String getView() {
        return view;
    }

    public void setView(String view) {
        this.view = view;
    }

    public Boolean getPredator() {
        return predator;
    }

    public void setPredator(Boolean predator) {
        this.predator = predator;
    }
}
