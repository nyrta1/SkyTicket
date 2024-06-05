package com.nyrta1.SkyTicket_AdminService.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.jdbc.datasource.DriverManagerDataSource;

import javax.sql.DataSource;

@Configuration
public class SupabasePostgreConfig {
    @Value("${spring.datasource.url}")
    private String supabasePostgreUrl;

    @Value("${spring.datasource.username}")
    private String supabasePostgreUsername;

    @Value("${spring.datasource.password}")
    private String supabasePostgrePassword;

    @Bean
    public DataSource dataSource() {
        DriverManagerDataSource dataSource = new DriverManagerDataSource();
        dataSource.setDriverClassName("org.postgresql.Driver");
        dataSource.setUrl(supabasePostgreUrl);
        dataSource.setUsername(supabasePostgreUsername);
        dataSource.setPassword(supabasePostgrePassword);
        return dataSource;
    }
}
