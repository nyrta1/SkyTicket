package com.nyrta1.SkyTicket_AdminService.repository;

import com.nyrta1.SkyTicket_AdminService.models.CountryEntity;

import java.util.List;

public interface SqlCrudOperation<T> {
    void add(T entity);
    T getById(int id);
    List<T> getAll();
    void updateById(int id, T entity);
    void deleteById(int id);
}