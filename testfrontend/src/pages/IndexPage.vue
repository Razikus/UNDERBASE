<template>
  <q-page class="flex flex-center">
    <q-card class="q-pa-md" style="width: 300px">
      <q-card-section>
        <div class="text-h6">Create User</div>
      </q-card-section>

      <q-card-section>
        <q-form @submit="onSubmit" class="q-gutter-md">
          <q-input
            v-model="email"
            label="Email"
            type="email"
            filled
            :rules="[val => !!val || 'Email is required', isValidEmail]"
          />

          <q-input
            v-model="password"
            label="Password"
            type="password"
            filled
            :rules="[val => !!val || 'Password is required', val => val.length >= 8 || 'Password must be at least 8 characters']"
          />

          <div>
            <q-btn label="Create User" type="submit" color="primary" />
          </div>
          <div >
            <q-btn label="Login" @click="loginNow" color="primary" />
          </div>
          <div >
            <q-btn :disable="token == ''" label="Execute" @click="executeNow" color="primary" />
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup>
import { ref } from 'vue';
import { useQuasar } from 'quasar';
import { authClient } from 'src/boot/auth';

defineOptions({
  name: 'IndexPage'
});

const $q = useQuasar();
const email = ref('test@test.com');
const password = ref('supersecretpassword!@#1');
const token = ref('');

const isValidEmail = (val) => {
  const emailPattern = /^(?=[a-zA-Z0-9@._%+-]{6,254}$)[a-zA-Z0-9._%+-]{1,64}@(?:[a-zA-Z0-9-]{1,63}\.){1,8}[a-zA-Z]{2,63}$/;
  return emailPattern.test(val) || 'Invalid email';
};

const loginNow = async () => {
  let res = await authClient.signInWithPassword({
    email: email.value,
    password: password.value,
  });

  if (res.error) {
    $q.notify({
      color: 'negative',
      message: 'Failed to login: ' + res.error.message,
      icon: 'warning'
    });
    return;
  } else {
    let mess = "WELCOME " + res.data.user.email
    token.value = res.data.session.access_token;
    $q.notify({
      color: 'positive',
      message: mess,
      icon: 'check'
    });
  }
};

const onSubmit = async () => {
  try {
    let res = await authClient.signUp({
      email: email.value,
      password: password.value,
    });
    if (res.error) {
      $q.notify({
        color: 'negative',
        message: 'Failed to create user: ' + res.error.message,
        icon: 'warning'
      });
      return;
    } else {
      $q.notify({
        color: 'positive',
        message: JSON.stringify(res),
        icon: 'check'
      });
    }

  } catch (error) {
    $q.notify({
      color: 'negative',
      message: 'Failed to create user: ' + error.message,
      icon: 'warning'
    });
  }
};

const executeNow = async () => {
  let backUrl = process.env.BACKEND + "/api/v1/hello"
  let res = await fetch(backUrl, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + token.value
    },
  });
  let data = await res.json();
  $q.notify({
    color: 'positive',
    message: "Backend endpoint executed: " + JSON.stringify(data),
    icon: 'check'
  });


}
</script>