<template>
    <v-container grid-list-xs fill-height >
        <v-layout row wrap align-center justify-center>
            <v-flex xs5 >
                <v-alert
                    :value="alert"
                    type="success"
                    dismissible
                    >
                    Register successfully.
                </v-alert>
                <v-card class="text-xs-center">
                    <v-card-text>
                        <v-container grid-list-xs>
                            <v-layout row wrap>
                                <v-flex xs12 class="mb-3">
                                    <span class="text-xs-center title">LOGIN</span>
                                </v-flex>
                                <v-flex xs12 class="px-3">
                                    <v-text-field
                                        ref="username"
                                        :rules="[rules.required]"
                                        name="username"
                                        label="Username"
                                        type="text"
                                        hint="Type your username"
                                        v-model="auth.username"
                                    ></v-text-field>
                                </v-flex>
                                <v-flex xs12 class="px-3">
                                    <v-text-field
                                        ref="password"
                                        :rules="[rules.required]"
                                        name="password"
                                        label="Password"
                                        type="password"
                                        hint="Type your password"
                                        v-model="auth.password"
                                    ></v-text-field>
                                </v-flex>
                                <v-flex xs12 class="mt-4 px-3">
                                    <v-btn block color="primary" @click="validateLogin()">SUBMIT</v-btn>
                                </v-flex>
                                <v-flex xs12 class="mt-3">
                                    <router-link to="/register" style="text-decoration:none;">
                                        <span class="text-xs-center body-1">Hit me to register new account</span>
                                    </router-link>
                                </v-flex>
                            </v-layout>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
export default {
    data(){
        return {
            auth: {
                username: '',
                password: ''
            },
            alert: false,
            rules: {
                required: val => !!val || 'This field is Required.'
            },
            hasError: false
        }
    },
    mounted(){
        if(this.$route.query.success){
            this.alert = true
        }
    },
    methods: {
        validateLogin(){
            this.hasError = false
            Object.keys(this.auth).forEach(v => {
                if(!this.auth[v]) this.hasError = true

                this.$refs[v].validate(true)
            })

            if(!this.hasError) this.login()
        },
        login(){
            this.$http.post('/user/login', this.auth)
            .then(resp => {
                if(resp && resp.status === 200){
                    this.$store.commit('setLoggedIn', true)
                    this.$store.commit('setUser', resp.data)
                    this.$router.push({ name: 'monitoring' })
                }
            })
        }
    }
}
</script>
