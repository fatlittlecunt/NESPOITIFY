from faker import Faker
from mysql.connector import connect, Error



try:
    with connect(
        host="localhost",
        user="root",
        password="root",
        database="MusicService"
    ) as connection:
        query="describe User"
        c = connection.cursor()
        c.execute(query)
        res = c.fetchall()
        for row in res:
            print(row)
except Error as e:
    print(e)


