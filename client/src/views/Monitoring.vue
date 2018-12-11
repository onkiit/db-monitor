<template>
  <v-container fluid grid-list-xs >
    <v-layout row wrap align-content-center justify-center>
      <v-flex xs4 xl4 class="">
        <psql-card/>
      </v-flex>
      <v-flex xs4 xl4 class="ml-1">
        <redis-card/>
      </v-flex>
      <v-flex xs4 xl4>
        <mongo-card/>
      </v-flex>
      <v-flex xs4 xl4 class="ml-1">
        <mysql-card />
      </v-flex>
    </v-layout>
    <v-dialog
      v-model="$store.state.modal"
      :overlay="false"
      max-width="500px"
      persistent
      transition="dialog-transition">
      <v-card class="pa-2">
        <v-card-title primary-title class="title">
          {{ $store.getters.modal_title }}
          <v-spacer></v-spacer>
          <v-icon @click="$store.commit('setModal', false)">close</v-icon>
        </v-card-title>
        <v-card-text>
          <div v-if="caller=='redis'">
            <redis-modal></redis-modal>
          </div>
          <div v-else-if="caller=='postgres'">
            <psql-modal></psql-modal>
          </div>
          <div v-else-if="caller=='mongo'">
            <mongo-modal></mongo-modal>
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn @click="$store.commit('setModal', false)" color="primary">close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
  import PsqlCard from '../components/card/PsqlCard'
  import MysqlCard from '../components/card/MysqlCard'
  import MongoCard from '../components/card/MongoCard'
  import RedisCard from '../components/card/RedisCard'
  
  import RedisModal from '../components/modal_content/RedisModal'
  import PsqlModal from '../components/modal_content/PsqlModal'
  import MongoModal from '../components/modal_content/MongoModal'
  export default {
    components: {
      PsqlCard,
      MysqlCard,
      MongoCard,
      RedisCard,

      RedisModal,
      PsqlModal,
      MongoModal
    },
    computed: {
      caller(){
        return this.$store.getters.modal_caller
      }
    }
  }
</script>