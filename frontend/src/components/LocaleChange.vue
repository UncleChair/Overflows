<script>
export default {
    methods: {
      changeLocale (locale) {
        this.$vuetify.locale.current = locale
        localStorage.setItem('locale', locale)
      },
    },
}
</script>
<template>
    <v-menu location="bottom">
        <template v-slot:activator="{ props }">
            <v-icon icon="$translate" v-bind="props" />
        </template>
        <v-list>
            <v-list-item title="简体中文" key="zh-CN" @click="changeLocale('zhHans')" />
            <v-list-item title="English" key="en" @click="changeLocale('en')" />
            <v-list-item :title="$t('system.moreLanguage')" key="more" disabled/>
        </v-list>
    </v-menu>
</template>