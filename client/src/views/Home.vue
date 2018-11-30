<template>
  <v-container fluid grid-list-xs >
    <v-layout row wrap align-content-center justify-center>
      <v-flex xs4 xl4 class="">
        <psql-card/>
      </v-flex>
      <v-flex xs4 xl4 class="ml-1">
        <redis-card />
      </v-flex>
      <v-flex xs4 xl4>
        <mongo-card/>
      </v-flex>
      <v-flex xs4 xl4 class="ml-1">
        <mysql-card />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  import PsqlCard from '../components/PsqlCard'
  import MysqlCard from '../components/MysqlCard'
  import MongoCard from '../components/MongoCard'
  import RedisCard from '../components/RedisCard'
  export default {
    components: {
      PsqlCard,
      MysqlCard,
      MongoCard,
      RedisCard
    }
  }
</script>