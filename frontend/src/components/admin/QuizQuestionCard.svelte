<script lang="ts">
    import type { QuizQuestion } from "../../model/quiz";
    import QuizChoiceCard from "./QuizChoiceCard.svelte";

    export let question: QuizQuestion;
</script>
<input type="text" placeholder="Question">
<div class="ml-2 mt-2 flex flex-col gap-2">
    {#each question.choices as choice(choice.id)}
        <QuizChoiceCard {choice} />
    {/each}
</div>