<template>
  <form>
    <label for="Username">Username:</label>
    <input v-model="name" name="Username" type="text" />

    <label for="Password">Password:</label>
    <input v-model="password" name="Password" type="text" />

    <button type="submit" @click="handleLogIn">Submit</button>

    {{ token }}
    {{ test }}
  </form>
</template>

<script setup lang="ts">
const name = ref("");
const password = ref("");

const token = useCookie("token")

const handleLogIn = async (e: MouseEvent) => {
  e.preventDefault();

  logIn({
    name: name.value,
    password: password.value,
  });
};
const test = ref<unknown>("a");
const { logIn } = useAuth();

onMounted(async () => {
  const { data } = await useApi("/user/3", {
 
  });


  test.value = data;
});
</script>

<style lang="css" scoped>
.test {
  width: 800px;
  background-color: red;
}
</style>
