package com.nyrta1.SkyTicket_AdminService.repository;

import com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class ManufacturerRepository implements SqlCrudOperation<ManufacturerEntity.Manufacturer> {
    @Autowired
    private JdbcTemplate jdbcTemplate;

    @Override
    public void add(ManufacturerEntity.Manufacturer entity) {
        String sqlQuery = "INSERT INTO manufacturer (name) VALUES (?)";
        jdbcTemplate.update(sqlQuery, entity.getName());
    }

    @Override
    public ManufacturerEntity.Manufacturer getById(int id) {
        String sqlQuery = "SELECT * FROM manufacturer WHERE id=?";
        return jdbcTemplate.queryForObject(sqlQuery, new Object[]{id}, (resultSet, i) -> {
            return ManufacturerEntity.Manufacturer.newBuilder()
                    .setId(resultSet.getInt("id"))
                    .setName(resultSet.getString("name"))
                    .build();
        });
    }

    @Override
    public List<ManufacturerEntity.Manufacturer> getAll() {
        String sqlQuery = "SELECT * FROM manufacturer";
        return jdbcTemplate.query(sqlQuery, (resultSet, i) -> {
            return ManufacturerEntity.Manufacturer.newBuilder()
                    .setId(resultSet.getInt("id"))
                    .setName(resultSet.getString("name"))
                    .build();
        });
    }

    @Override
    public void updateById(int id, ManufacturerEntity.Manufacturer entity) {
        String sqlQuery = "UPDATE manufacturer SET name=? WHERE id=?";
        jdbcTemplate.update(sqlQuery, entity.getName(), id);
    }

    @Override
    public void deleteById(int id) {
        String sqlQuery = "DELETE FROM manufacturer WHERE id=?";
        jdbcTemplate.update(sqlQuery, id);
    }
}
