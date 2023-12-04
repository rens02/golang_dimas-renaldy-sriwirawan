- Gambarkan desain ERD dari sistem pembelian tiket kereta api!

![erd_tiket_kereta](erd_tiket_kereta.png)

- Gambarkan use case diagram dari sistem pembelian tiket kereta api!

![usecase_tiket_kereta](usecase_tiket_kereta.png)

- Terdapat sebuah query pada SQL yaitu `SELECT * FROM users;` Dengan tujuan yang sama, tuliskan dalam bentuk perintah:
    - Redis

    ```
    HGETALL users
    ```

    - Neo4j
    
    ```
    MATCH (u:User)
    RETURN u;
    ```
    
    - Cassandra

    ```
    SELECT * FROM users;
    ```