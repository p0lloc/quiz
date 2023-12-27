<script lang="ts">
    import type { Quiz } from "../../model/quiz";

    export let quiz: Quiz;

    let code: string | null = null;

    async function host(){
        let response = await fetch(
            `http://localhost:3000/api/quizzes/${quiz.id}/host`, {
                method: "post"
        });

        let json = await response.json();
        code = json.code;
    }
</script>

<div class="border rounded-md p-4 flex justify-between items-center">
    <p class="font-bold">{quiz.name} - {quiz.questions.length} questions </p>
    <div>
        {#if code != null}
            {code}
        {:else}
            <button on:click={host}>Host</button>
        {/if}
    </div>
</div>