package com.nyrta1.SkyTicket_AdminService.repository;

import com.nyrta1.SkyTicket_AdminService.models.AirplaneEntity;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.sql.Timestamp;
import java.time.Instant;
import java.util.List;

@Repository
public class AirplaneRepository implements SqlCrudOperation<AirplaneEntity.Airplane> {
    @Autowired
    private JdbcTemplate jdbcTemplate;

    @Override
    public void add(AirplaneEntity.Airplane airplane) {
        String sqlQuery = "INSERT INTO airplane (manufacturer_id, manufacturer_year, first_slot_capacity, " +
                "economy_slot_capacity, country_origin_id, created_at) " +
                "VALUES (?, ?, ?, ?, ?, ?)";
        jdbcTemplate.update(sqlQuery, airplane.getManufacturerId(), airplane.getManufacturerYear(),
                airplane.getFirstSlotCapacity(), airplane.getEconomySlotCapacity(),
                airplane.getCountryOriginId(), new Timestamp(System.currentTimeMillis()));
    }

    @Override
    public AirplaneEntity.Airplane getById(int id) {
        String sqlQuery = "SELECT * FROM airplane WHERE id=?";
        return jdbcTemplate.queryForObject(sqlQuery, new Object[]{id}, (resultSet, i) -> {
            return AirplaneEntity.Airplane.newBuilder()
                    .setId(resultSet.getLong("id"))
                    .setManufacturerId(resultSet.getShort("manufacturer_id"))
                    .setManufacturerYear(resultSet.getShort("manufacturer_year"))
                    .setFirstSlotCapacity(resultSet.getShort("first_slot_capacity"))
                    .setEconomySlotCapacity(resultSet.getShort("economy_slot_capacity"))
                    .setCountryOriginId(resultSet.getShort("country_origin_id"))
                    .build();
        });
    }

    @Override
    public List<AirplaneEntity.Airplane> getAll() {
        String sqlQuery = "SELECT * FROM airplane";
        return jdbcTemplate.query(sqlQuery, (resultSet, i) -> {
            return AirplaneEntity.Airplane.newBuilder()
                    .setId(resultSet.getLong("id"))
                    .setManufacturerId(resultSet.getShort("manufacturer_id"))
                    .setManufacturerYear(resultSet.getShort("manufacturer_year"))
                    .setFirstSlotCapacity(resultSet.getShort("first_slot_capacity"))
                    .setEconomySlotCapacity(resultSet.getShort("economy_slot_capacity"))
                    .setCountryOriginId(resultSet.getShort("country_origin_id"))
                    .build();
        });
    }

    @Override
    public void updateById(int id, AirplaneEntity.Airplane airplane) {
        String sqlQuery = "UPDATE airplane SET manufacturer_id=?, manufacturer_year=?, " +
                "first_slot_capacity=?, economy_slot_capacity=?, country_origin_id=?, updated_at=? " +
                "WHERE id=?";
        jdbcTemplate.update(sqlQuery, airplane.getManufacturerId(), airplane.getManufacturerYear(),
                airplane.getFirstSlotCapacity(), airplane.getEconomySlotCapacity(),
                airplane.getCountryOriginId(), new Timestamp(System.currentTimeMillis()), id);
    }

    @Override
    public void deleteById(int id) {
        String sqlQuery = "DELETE FROM airplane WHERE id=?";
        jdbcTemplate.update(sqlQuery, id);
    }
}
