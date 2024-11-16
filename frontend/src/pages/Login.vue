<route>
    {
        meta: {
            layout: "empty",
        }
    }
</route>
<script>
import { useUserInfoStore } from "@/stores/userInfo";
import axios from "axios";
import { mapStores } from "pinia";

export default {
    data() {
        return {
            user: null,
            password: null,
            loading: false,
            showPassword: false,
            form: "",
            errorMessage: null,
        }
    },
    computed: {
        ...mapStores(useUserInfoStore)
    },
    methods: {
        isEmail(str) {
            var reg = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
            return reg.test(str);
        },
        async login() {
            let email = null;
            let username = null;
            if (this.isEmail(this.user)) {
                email = this.user;
            }
            else {
                username = this.user;
            }
            this.loading = true;
            try {
                let loginResponse = await axios.post("/auth/login", {
                    username: username,
                    email: email,
                    password: this.password
                })
                this.errorMessage = null;
                this.userInfoStore.jwt = loginResponse.data.token;
                this.userInfoStore.jwtExpiry = loginResponse.data.expire;
            }
            catch (error) {
                console.error(error)
                if (error.response == null) {
                    this.errorMessage = this.$t('system.networkError')
                } else {
                    this.errorMessage = this.$t('login.loginFailed')
                }
                this.loading = false;
                return
            }
            // get user info
            try {
                let userInfoResponse = await axios.get("/users/current")
                this.userInfoStore.uid = userInfoResponse.data.uid;
                this.userInfoStore.username = userInfoResponse.data.username;
                this.userInfoStore.email = userInfoResponse.data.email;
                this.userInfoStore.avatarUrl = userInfoResponse.data.avatarUrl;
            }
            catch (error) {
                console.error(error)
            }
            this.loading = false;
            this.$router.push("/user");
        },
        onSubmit() {
            if (!this.form) return
            this.loading = true
            this.login();
        },
        required(v) {
            return !!v || this.$t('system.required')
        },
    }
}
</script>
<template>
    <v-container class="h-100 pa-0" fluid>
        <v-row align="center" class="h-100" justify="center">
            <v-responsive class="flex-1-1 px-4" max-width="475">
                <v-img class="mx-auto mb-4" max-width="60" src="@/assets/logo.svg" />
                <div class="text-h5 text-center mb-8 font-weight-medium">
                    {{ $t('login.title') }}
                </div>
                <v-card class="pa-10 mb-8" elevation="3" rounded="lg">
                    <v-form v-model="form" @submit.prevent="onSubmit">
                        <v-label class="text-subtitle-2">{{ $t('login.usernameOrEmail') }}</v-label>
                        <v-text-field v-model="this.user" :readonly="loading" :rules="[required]" color="primary"
                            density="compact" rounded="lg" variant="outlined" />
                        <v-label class="text-subtitle-2">{{ $t('login.password') }}</v-label>
                        <v-text-field v-model="this.password" :readonly="loading" :rules="[required]"
                            :type="this.showPassword ? 'text' : 'password'" color="primary"
                            @click:append-inner="this.showPassword = !this.showPassword"
                            :append-inner-icon="this.showPassword ? '$eye' : '$eyeOff'" density="compact" rounded="lg"
                            variant="outlined" />
                        <div class="mb-4">
                            <div class="d-flex justify-space-between align-center">
                                <v-checkbox-btn v-model="userInfoStore.rememberMe" class="ms-n3" color="primary">
                                    <template #label>
                                        <span class="text-body-2">{{ $t('login.rememberMe') }}</span>
                                    </template>
                                </v-checkbox-btn>
                                <a class="text-decoration-none text-primary text-body-2 font-weight-medium" href="#">
                                    {{ $t('login.forgetPassword') }}
                                </a>
                            </div>
                        </div>
                        <v-btn :disabled="!form" :loading="loading" type="submit" block class="text-none"
                            color="primary" flat rounded="lg" :text="$t('login.login')" />
                        <v-messages color="error" :active="errorMessage !== null"
                            class="text-subtitle-1 pt-2 d-flex justify-center" :messages="errorMessage"></v-messages>
                    </v-form>
                </v-card>
                <div class="text-center text-body-2">
                    {{ $t('login.noAccount') }}
                    <router-link class="text-decoration-none text-primary font-weight-medium"
                        :to="{ path: '/signup' }">{{
                            $t('login.signup') }}</router-link>
                </div>
            </v-responsive>
        </v-row>
    </v-container>
</template>