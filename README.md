# local-ai

This is only a docker container to pull build a local Ollama chat engine.
Go to your docker folder
```
cd docker
docker-compose up -d
```

In you browser:
```
http://localhost:3000
```

- Sign up, the first signing up user will be the admin user.
- Locate settings
- Look for models
- Download model like llama3
- Go back to the main chat window, select the downloaded model and enjoy your local chat engine


Also why don't you try generations images? There is a complete ready image for this:

```
git@github.com:AbdBarho/stable-diffusion-webui-docker.git

cd stable-diffusion-webui-docker

docker compose --profile download up --build
# wait until its done, then:
docker compose --profile [ui] up --build
# where [ui] is one of: auto | auto-cpu | comfy | comfy-cpu

```

Visit: http://localhost:7860/

It's really fun. Try it

You can find an example go implementation in cmd folder to ask question from llama3 model as a command line tool:
