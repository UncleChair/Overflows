<template>
  <v-app>
    <v-app-bar color="transparent" flat>
      <template #prepend>
        <v-app-bar-nav-icon v-if="$vuetify.display.smAndDown" @click="drawer = !drawer" />
      </template>

      <div class="d-md-flex d-none ga-4 mx-auto">
        <v-btn v-for="item in appBarItems" :key="item" class="text-none" :text="item" />
      </div>

      <template #append>
        <v-app-bar-nav-icon class="ms-auto text-none" slim>
          <ThemeChange />
        </v-app-bar-nav-icon>
        <v-app-bar-nav-icon class="ms-auto text-none" slim>
          <LocaleChange />
        </v-app-bar-nav-icon>
        <template v-if="this.logged">
          <v-app-bar-nav-icon class="ms-auto text-none" slim>
            <UserAvatar />
          </v-app-bar-nav-icon>
        </template>
        <template v-else>
          <v-btn append-icon="$chevronRight" class="ms-auto text-none" slim text="Login"
            @click="() => { $router.push({ path: '/login' }) }" />
        </template>
      </template>
    </v-app-bar>
    <v-main>
      <router-view />
    </v-main>

    <AppFooter />
  </v-app>
</template>

<script setup>
//
</script>
