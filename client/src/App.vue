<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>dbcheck</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn flat medium fab @click="refresh()">
        <v-icon>refresh</v-icon>
      </v-btn>
    </v-toolbar>
    <v-content>
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  methods: {
    refresh(){
      window.location.reload()
    }
  }
}
</script>