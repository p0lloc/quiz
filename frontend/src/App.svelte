<script>
    import JoinGame from "./view/JoinGame.svelte";
    import Login from "./view/Login.svelte";
    import QuizList from "./view/QuizList.svelte";
    import Router from 'svelte-spa-router'
    import Register from "./view/Register.svelte";

    const routes = {
        '/': JoinGame,
        '/list': QuizList,
        '/login': Login,
        '/register': Register,
    }
</script>

<Router routes={routes} />