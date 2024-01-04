<script lang="ts">
    let form: HTMLFormElement;

    async function login() {
        if(!form.reportValidity())
            return;

        const formData = new FormData(form);
        let response = await fetch(`http://localhost:3000/auth/login`, {
            method: "post",
            body: formData
        });

    }
</script>

<div class="main min-h-screen min-w-full bg-slate-400 flex justify-center items-center">
    <div class="bg-white w-1/2 rounded-xl p-8">
        <h2 class="text-3xl font-bold">Login to Quiz</h2>
        <form bind:this={form}>
            <input name="username" type="email" placeholder="E-mail" class="mt-4 w-full">
            <input name="password" type="password" placeholder="Password" class="mt-2 w-full">
        </form>
        <div class="flex justify-end mt-4">
            <button on:click={login}>Log in</button>
        </div>
    </div>
</div>