<route>
    {
        meta: {
            layout: "empty",
        }
    }
</route>
<script>
import { useUserInfoStore } from "@/stores/userInfo";
import { mapStores } from "pinia";
import axios from "axios";

export default {
    data() {
        return {
            username: null,
            email: null,
            password: null,
            loading: false,
            showPassword: false,
            form: "",
            errorProcessingDialog: false,
            error: null,
            errorTitle: null,
            errorMessage: null,
        }
    },
    computed: {
        ...mapStores(useUserInfoStore)
    },
    methods: {
        isEmail(v) {
            var reg = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
            return reg.test(v) || this.$t('system.invalidEmail')
        },
        isPassword2(v) {
            var reg = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).{6,18}$/;
            return reg.test(v) || this.$t('system.invalidPassword')
        },
        signup() {
            axios.post("/auth/signup", {
                user_name: this.username,
                email: this.email,
                password: this.password,
            }).then((response) => {
                this.userInfoStore.jwt = response.data.token;
                this.loading = false;
                this.$router.push("/user");
            }).catch((error) => {
                this.error = error;
                this.errorProcessingDialog = true;
                if (error.response.status === 409) {
                    if (error.response.data.code == 1) {
                        this.errorTitle = this.$t('signup.errorTitle');
                        this.errorMessage = this.$t('signup.emailError');
                    }
                    else if (error.response.data.code == 2) {
                        this.errorTitle = this.$t('signup.errorTitle');
                        this.errorMessage = this.$t('signup.usernameError');
                    }
                }
                else {
                    this.errorTitle = null;
                    this.errorMessage = null;
                }
                this.loading = false;
            });
        },
        onSubmit() {
            if (!this.form) return
            this.loading = true
            this.signup();
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
                    {{ $t('signup.title') }}
                </div>
                <v-card class="pa-10 mb-4" elevation="3" rounded="lg">
                    <v-form v-model="form" @submit.prevent="onSubmit">
                        <v-label class="text-subtitle-2">{{ $t('signup.username') }}</v-label>
                        <v-text-field v-model="this.username" :rules="[required]" color="primary" density="compact"
                            rounded="lg" variant="outlined" />
                        <v-label class="text-subtitle-2">{{ $t('signup.email') }}</v-label>
                        <v-text-field v-model="this.email" :rules="[required, isEmail]" color="primary"
                            density="compact" rounded="lg" variant="outlined" />
                        <v-label class="text-subtitle-2">{{ $t('signup.password') }}</v-label>
                        <v-text-field v-model="this.password" :rules="[required, isPassword2]"
                            :type="this.showPassword ? 'text' : 'password'" color="primary"
                            @click:append-inner="this.showPassword = !this.showPassword"
                            :append-inner-icon="this.showPassword ? '$eye' : '$eyeOff'" density="compact" rounded="lg"
                            variant="outlined" />
                        <v-btn block type="submit" class="text-none mt-4" color="primary" rounded="lg"
                            :text="$t('signup.signup')" />
                    </v-form>
                </v-card>
                <div class="text-center text-body-2">
                    {{ $t('signup.hasAccount') }}
                    <router-link class="text-decoration-none text-primary font-weight-medium"
                        :to="{ path: '/login' }">{{
                            $t('signup.login') }}</router-link>
                </div>
            </v-responsive>
        </v-row>
    </v-container>
</template>