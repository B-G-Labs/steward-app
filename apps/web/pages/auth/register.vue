<template>
  <form>
    <label for="Username">Username:</label>
    <input v-model="name" name="Username" type="text" />

    <label for="Password">Password:</label>
    <input v-model="password" name="Password" type="text" />

    <button type="submit" @click="handleRegister">Submit</button>
  </form>
</template>

<script setup lang="ts">
const name = ref("");
const password = ref("");

const handleRegister = async (e: MouseEvent) => {
  e.preventDefault()

  register({
    name: name.value,
    password: password.value
  })
}

const { register } = useAuth();
</script>

<style lang="css" scoped>
.test {
  width: 800px;
  background-color: red;
}
</style>
