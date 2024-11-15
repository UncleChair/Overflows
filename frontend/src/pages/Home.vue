<template>
  <v-container class="fill-height">
    <v-responsive class="align-centerfill-height mx-auto" max-width="900">
      <v-img class="mb-4" height="150" src="@/assets/logo.png" />
      <div class="text-center">
        <div class="text-body-2 font-weight-light mb-n1">Welcome to</div>
        <h1 class="text-h2 font-weight-bold">Overflows</h1>
      </div>
      <div class="py-4" />
      <div class="d-flex ga-4 justify-center">
        <v-btn class="text-none" color="primary" flat rounded="lg" :text="$t('homePage.getStarted')"
          @click="() => { $router.push({ path: '/user' }) }" />
        <v-btn class="text-none" flat rounded="lg" :text="$t('homePage.learnUsage')"
          @click="() => { $router.push({ path: '/guide' }) }" />
      </div>
    </v-responsive>
  </v-container>
</template>

<script setup>
</script>
