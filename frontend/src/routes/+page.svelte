<script>
    import { onMount } from "svelte";

    let username = '';

    let count = 0;

    onMount(async function() {
        username = payload.username;
    });
</script>

<svelte:head>
    <title>Hello {username}</title>
</svelte:head>

<main>
    <h1>Welcome</h1>
    <p>Your username is {username}</p>

    <input type="button" value="You have pressed this button {count} times." on:click={() => count++}>
</main>
