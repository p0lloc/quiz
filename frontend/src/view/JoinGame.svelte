<script lang="ts">
    let code = 123456;

    async function join(){
        let response = await fetch(`http://localhost:3000/join?code=${code}`, {
           method: "post"
        });

        if(!response.ok)
            return;

        setupWs();
    }

    function setupWs(){
        const socket = new WebSocket("ws://localhost:3000/ws");

        // Connection opened
        socket.addEventListener("open", (event) => {
            setInterval(() => {
                socket.send("Hello Server!");
            }, 1000);
        });

        // Listen for messages
        socket.addEventListener("message", (event) => {
            console.log("Message from server ", event.data);
        });
    }
</script>

<div class="main">
    <h1>Quiz</h1>
    <div id="input">
        <input type="number" placeholder="Game code" bind:value={code} />
        <button on:click={join}>Join</button>
    </div>
</div>

<style>
    h1 {
        color: white;
        text-align: center;
        font-size: 72px;
    }

    #input {
        margin-top: 40px;
    }

    input, button {
        padding: 10px 20px;
    }

    .main {
        padding-top: 10rem;
        background-color: darkorchid;
        display: flex;
        flex-direction: column;
        align-items: center;
        min-height: 100vh;
        min-width: 100vw;
    }
</style>