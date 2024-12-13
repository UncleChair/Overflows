import request from "@/request/instance";
import { defineStore } from "pinia";

export const useUserInfoStore = defineStore("userInfo", {
    state: () => ({
        uid: "",
        username: "",
        email: "",
        avatarUrl: "",
        jwt: "",
        jwtExpiry: 0,
        sessionId: "",
        rememberMe: false,
    }),
    getters: {
        getUserInfo: (state) => state,
    },
    actions: {
        async syncUserBasicInfo() {
            await request.get("/users/current")
                .then((response) => {
                    this.uid = response.data.uid;
                    this.username = response.data.username;
                    this.email = response.data.email;
                    this.avatarUrl = response.data.avatarUrl;
                })
                .catch((error) => {
                    const errorCode = error.code;
                    const errorMessage = error.message;
                    console.error('Sync user basic info failed: ' + errorCode + ' ' + errorMessage);
                });
        },
        async getUserJwt() {
            // TODO: auto extend jwt if it's about to expire or expired in 24 hours

            return this.jwt;
        }
    },
    // Use local storage for user info
    persist: true,
})