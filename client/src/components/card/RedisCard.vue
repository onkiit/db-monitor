<template>
    <v-card height="250">
        <v-card-title>
            <span class="title">Redis</span>
            <v-spacer></v-spacer>
            <v-icon @click="$store.getters['redis/disabled'] ? '': modal()" :color="$store.getters['redis/disabled'] ? 'grey lighten-2':'primary'">info</v-icon>
        </v-card-title>
        <v-card-text>
            <v-layout v-if="$store.getters['redis/disabled']" row wrap>
                <span class="body-2">Redis is disabled by server.</span>
            </v-layout>
           <v-layout v-else row wrap>
                <v-flex xs12>
                    <v-flex xs6>
                        <span class="body-2">Version</span>
                    </v-flex>
                    <v-flex xs6>
                        <span class="caption">{{ version }}</span>
                    </v-flex>
                </v-flex>
                <v-flex xs12>
                    <v-flex xs6>
                        <span class="body-2">Active Clients</span>
                    </v-flex>
                    <v-flex xs6>
                        <span class="caption"><strong>{{ active_client }}</strong> {{active_client > 2 ? "Active clients": "Active client"}}</span>
                    </v-flex>
                </v-flex>
            </v-layout>
        </v-card-text>
    </v-card>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from 'vuex'
export default {
    created(){
        this.setVersion
        this.setActiveClient
        this.setHealth
    },
    computed: {
        ...mapGetters({
            version: 'redis/version',
            active_client: 'redis/active_client',
            health: 'redis/health'
        }),
        ...mapActions({
            setVersion: 'redis/setVersion',
            setActiveClient: 'redis/setActiveClient',
            setHealth: 'redis/setHealth'
        })
    },
    methods: {
        ...mapMutations({
            setModal: 'setModal',
            setModalTitle: 'setModalTitle',
            setModalCaller: 'setModalCaller',
        }),
        modal(){
            this.setModal(true)
            this.setModalTitle('Redis Health Information')
            this.setModalCaller('redis')
        }
    }
}
</script>