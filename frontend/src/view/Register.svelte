<script lang="ts">
    let form: HTMLFormElement;

    async function register() {
        if(!form.reportValidity())
            return;

        const formData = new FormData(form);
        let response = await fetch(`http://localhost:3000/auth/register`, {
            method: "post",
            body: formData
        });

    }
</script>

<div class="main min-h-screen min-w-full bg-slate-400 flex justify-center items-center">
    <div class="bg-white w-1/2 rounded-xl p-8">
        <h2 class="text-3xl font-bold">Register to Quiz</h2>
        <form bind:this={form}>
            <input name="username" type="email" placeholder="E-mail" class="mt-4 w-full">
            <input name="password" type="password" placeholder="Password" class="mt-4 w-full">
            <input name="confirm" type="password" placeholder="Confirm password" class="mt-2 w-full">
        </form>
        <div class="flex justify-end mt-4">
            <button on:click={register}>Sign up</button>
        </div>
    </div>
</div>