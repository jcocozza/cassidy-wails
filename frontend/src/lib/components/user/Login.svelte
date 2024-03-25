<script lang="ts">
    import { goto } from "$app/navigation";
    import type { ActionData } from '../../../routes/(unprotected)/auth/login/$types';
    export let form: ActionData;

    let username: string;
    let password: string;
    let invalid_password: boolean = false;

    function toSignUp() {
      goto("/auth/signup")
    }

    $: {
      if (password) {
        invalid_password = false;
      }
    }
</script>

<div class="login container">
  <h3>Login</h3>
  <form method="POST" action="?/login">
    <input class="form-control" bind:value="{username}" type="email" name="username" placeholder="Email" required />
    <input class="form-control" bind:value="{password}" type="password" name="password" placeholder="Password" required />
    {#if form?.success == "incorrect password"  }
      <p style="color: red;">Incorrect password!</p>
    {/if}

    <input type="submit">
  </form>
  <p>
    Don't have an account?
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
    <strong class="pe-auto" on:click={toSignUp}>Sign up</strong>
  </p>
</div>

