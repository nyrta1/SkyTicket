package com.nyrta1.SkyTicket_AdminService.repository;

import com.nyrta1.SkyTicket_AdminService.models.CountryEntity;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class CountryRepository implements SqlCrudOperation<CountryEntity.Country> {
    @Autowired
    private JdbcTemplate jdbcTemplate;

    @Override
    public void add(CountryEntity.Country country) {
        String sqlQuery = "INSERT INTO country (name, prefix) VALUES (?, ?)";
        jdbcTemplate.update(sqlQuery, country.getName(), country.getPrefix());
    }

    @Override
    public CountryEntity.Country getById(int id) {
        String sqlQuery = "SELECT * FROM country WHERE id=?";
        return jdbcTemplate.queryForObject(sqlQuery, new Object[]{id}, (resultSet, i) -> {
            return CountryEntity.Country.newBuilder()
                    .setId(resultSet.getInt("id"))
                    .setName(resultSet.getString("name"))
                    .setPrefix(resultSet.getString("prefix"))
                    .build();
        });
    }

    @Override
    public List<CountryEntity.Country> getAll() {
        String sqlQuery = "SELECT * FROM country";
        return jdbcTemplate.query(sqlQuery, (resultSet, i) -> {
            return CountryEntity.Country.newBuilder()
                    .setId(resultSet.getInt("id"))
                    .setName(resultSet.getString("name"))
                    .setPrefix(resultSet.getString("prefix"))
                    .build();
        });
    }

    @Override
    public void updateById(int id, CountryEntity.Country country) {
        String sqlQuery = "UPDATE country SET name=?, prefix=? WHERE id=?";
        jdbcTemplate.update(sqlQuery, country.getName(), country.getPrefix(), id);
    }

    @Override
    public void deleteById(int id) {
        String sqlQuery = "DELETE FROM country WHERE id=?";
        jdbcTemplate.update(sqlQuery, id);
    }
}
