<script lang="ts">
    import { goto } from "$app/navigation";
    import { AuthenticateUser } from "$lib/wailsjs/go/controllers/UserHandler";
    import { SetUser } from "$lib/wailsjs/go/main/App";
    import type { controllers } from "$lib/wailsjs/go/models";

    let invalid_password: boolean = false;
    let other_auth_error: boolean = false;
    let authRequest: controllers.authRequest = {
    	username: "",
    	password: ""
    };

    function toSignUp() {
    	goto("/auth/signup")
    }

    async function login() {
    	try {
        	invalid_password = false;
        	other_auth_error = false;
        	let usr = await AuthenticateUser(authRequest)
        	await SetUser(usr)
        	goto("/microcycle")
      } catch (error) {
    	console.error(error)
        if (error === "incorrect password") {
        	invalid_password = true;
        } else {
         	other_auth_error = true;
        }
      }
    }
</script>

<div class="login container">
	<h3>Login</h3>
	<form>
    	<input class="form-control" bind:value="{authRequest.username}" type="email" name="username" placeholder="Email" required/>
    	<input class="form-control" bind:value="{authRequest.password}" type="password" name="password" placeholder="Password" required/>
		{#if invalid_password}
			<p style="color: red;">Incorrect password!</p>
		{/if}
		{#if other_auth_error}
			<p style="color: red;">Authentication Error</p>
		{/if}
		<button class="btn btn-primary" on:click={async () => {await login()}}>Login</button>
  	</form>
  	<p>
    	Don't have an account?
    	<!-- svelte-ignore a11y-click-events-have-key-events -->
    	<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
 		<strong class="pe-auto" on:click={toSignUp}>Sign up</strong>
	</p>
</div>

