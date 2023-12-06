console.time('library import')
const mysql = require('mysql2/promise');
const { faker } = require('@faker-js/faker');
console.timeEnd('library import')

async function insertDummyUsers() {
  console.time('DB connection Establishment')
  const connection = await mysql.createConnection({
    host: 'mysql',
    user: 'root',
    password: 'root',
    database: 'social_network',
  });
  console.timeEnd('DB connection Establishment')

  console.time('dummyUsersGeneration');
  const dummyUsers = Array.from({ length: 10 }, () => [
    faker.internet.userName(),
    faker.internet.email(),
    faker.internet.password(),
    faker.image.url(),
  ]);
  console.timeEnd('dummyUsersGeneration');

  console.time('insertIntoDB');
  await connection.query('INSERT INTO users (username, email, password, profile_image_url) VALUES ?', [dummyUsers]);
  console.timeEnd('insertIntoDB');
  await connection.end();

}

console.time('totalExecutionTime');
insertDummyUsers().then(() => {
  console.timeEnd('totalExecutionTime');
});
