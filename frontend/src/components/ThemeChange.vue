<script>
import { useTheme } from 'vuetify'

export default {
    setup() {
        const theme = useTheme()
        return {
            theme,
            toggleTheme() {
                if (theme.global.current.value.dark) {
                    theme.global.name.value = 'light';
                    localStorage.setItem('theme', 'light');
                }
                else {
                    theme.global.name.value = 'dark';
                    localStorage.setItem('theme', 'dark');
                }
            }
        }
    },
}
</script>
<template>
    <template v-if="theme.global.name.value == 'light'">
        <v-icon icon="$weatherSunny" @click="toggleTheme" />
    </template>
    <template v-if="theme.global.name.value == 'dark'">
        <v-icon icon="$weatherNight" @click="toggleTheme" />
    </template>
</template>