// Contoh berkas Cypher untuk Neo4j

// Data Contoh (pengguna)
CREATE (:User {name: "tiar"})
CREATE (:User {name: "tiara"})
CREATE (:User {name: "tiari"})

// Ambil data dari pengguna
MATCH (u:User)
RETURN u