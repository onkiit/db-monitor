<template>
    <v-container grid-list-xs fill-height >
        <v-layout row wrap align-center justify-center>
            <v-flex xs5 >
                <v-card class="text-xs-center">
                    <v-card-text>
                        <v-container grid-list-xs>
                            <v-layout row wrap>
                                <v-flex xs12 class="mb-3">
                                    <span class="text-xs-center title">Register</span>
                                </v-flex>
                                <v-flex xs12 class="px-3">
                                    <v-text-field
                                        ref="name"
                                        :rules="[rules.required]"
                                        name="name"
                                        label="Name"
                                        type="text"
                                        hint="Type your Name"
                                        v-model="auth.name"
                                    ></v-text-field>
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
                                        ref="email"
                                        :rules="[rules.required]"
                                        name="email"
                                        label="Email"
                                        type="text"
                                        hint="Type your email"
                                        v-model="auth.email"
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
                                        append-icon="arrow_down"
                                    ></v-text-field>
                                </v-flex>
                                <v-flex xs12 class="mt-4 px-3">
                                    <v-btn block color="primary" @click="validateRegister()">SUBMIT</v-btn>
                                </v-flex>
                                <v-flex xs12 class="mt-3">
                                    <router-link to="/login" style="text-decoration:none;">
                                        <span class="text-xs-center body-1">Login to your existing account.</span>
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
                name: '',
                username: '',
                email: '',
                password: ''
            },
            rules: {
                required: val => !!val || 'This field is Required.',
                username: val => {
                    this.$http.get('/user/checkuser', { username: val })
                    .then(resp => {
                        if(resp && resp.data.is_registered){
                            return 'This username has already registered'
                        }
                    })
                }
            },
            hasError: false
        }
    },
    methods: {
        validateRegister(){
            this.hasError = false
            Object.keys(this.auth).forEach(v => {
                if(!this.auth[v]) this.hasError = true
                this.$refs[v].validate(true)
            })

            if(!this.hasError){
                this.register()
            }
        },
        register(){
            this.$http.post('/user/register', this.auth)
            .then(resp => {
                if(resp && resp.status === 200){
                    this.$router.push('/login?success=true')
                }
            })
        }
    }
}
</script>
