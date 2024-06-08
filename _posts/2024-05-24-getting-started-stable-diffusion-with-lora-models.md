---
title: Getting Started Stable Diffusion with LoRA models
tags:
  - lora
  - stable diffusion
  - python
  - automatic1111
toc: true
date: 2024/05/29
---

## 1. About Stable Diffusion

There are various terminologies and explanations about stable diffusion and LoRA, so it's better to read them first:
I'm also a beginner about these topics.

- [Hugging Face: Using LoRA for Efficient Stable Diffusion Fine-Tuning](https://huggingface.co/blog/lora)
- [Anakin.ai: How to Use Lora in Stable Diffusion - A Step-by-Step Guide](https://anakin.ai/blog/how-to-use-lora-stable-diffusion/)
- [Anakin.ai: How to Use Stable Diffusion Checkpoints - A Complete Guide](https://anakin.ai/blog/how-to-use-stable-diffusion-checkpoint/)

In this article, I followed [the article in Anakin.ai](https://anakin.ai/blog/how-to-use-lora-stable-diffusion/) mostly.


## 2. LoRA models

There are some web services in which we can download LoRA models.
For example, I downloaded the following models from [civitai](https://civitai.com/).

- [Detail Tweaker LoRA](https://civitai.com/models/58390/detail-tweaker-lora-lora)

In order to use the model, in this article, I explored setting up a stable diffusion web UI in WSL 2 by following [this manual](https://github.com/civitai/civitai/wiki/How-to-use-models).


## 3. Set up a web UI for stable diffusion in WSL 2

### 3.1. Install NVIDIA driver

First, follow [this guide](https://docs.nvidia.cn/cuda/wsl-user-guide/index.html#getting-started-with-cuda-on-wsl-2) to install an NVIDIA driver.

- Install an NVIDA driver on a Windows host machine
- Make sure WSL 2 is up to date by `wsl --update`

Then, just in case, I restarted my laptop.
I couldn't find any CUDA Toolkit on the download page for the WSL-Ubuntu page, so I skipped it.

### 3.2. Set up an automatic1111 web UI

Next, set up a web UI of AUTOMATIC1111, which is a web UI for stable diffusion, by following [this guide](https://github.com/AUTOMATIC1111/stable-diffusion-webui?tab=readme-ov-file#automatic-installation-on-linux).
Also, there is [another component](https://github.com/AUTOMATIC1111/stable-diffusion-webui/issues/10117#issuecomment-1536437860) that needs to be installed for WSL 2.

```bash
sudo apt install python3-venv libgl1 libglib2.0-0
sudo apt install --no-install-recommends google-perftools
sudo apt install gc # probably necessary
```

Then, run a web UI.
```bash
wget -q https://raw.githubusercontent.com/AUTOMATIC1111/stable-diffusion-webui/master/webui.sh
bash webui.sh
```

Then access to `http://127.0.0.1:7860` and see if a web UI can be seen.

## 4. Generate images on AUTOMATIC 1111

Here, I just focused on testing using Lora models and tested them.

### 4.1. How to load a Lora model in a web UI

Download a Lora model from civitai and put it under `stable-diffusion-webui/models/Lora/`.
Then, on a web UI, select a LoRA model, and the prompt adds a prefix like `<lora:add_detail:1>`.
The `1` in the prefix is a weight that allows us to change the influence of a LoRA model.


### 4.2. Set a checkpoint model

As a default, there is a `v1.5-pruned-emaonly.safetensors`.
This checkpoint is a stable diffusion v1.5 model, and it seems the Pruned, EMA model is suitable just to generate images, but not for training, according to [this article](https://stable-diffusion-art.com/models/#Pruned_vs_Full_vs_EMA-only_models).
Note that in Civitai, the Base Model of some LoRA models is `SD 1.5`, which is the stable diffusion v1.5 model. However, this doesn't mean that similar images shown in Civitai can be generated.
There is some information that was used to generate such as checkpoint images, Lora images, prompts, and negative prompts for some images in the Civitai.
For example, for the image for [this model](https://civitai.com/models/15743/lora-stacia-goddess-of-creation-or-asuna-or-underworld-or-sao), there is information like below, so we can see if the configurations and set up of AUTOMATIC1111 is correct.

![generated information](/assets/images/posts/2024/05/29/getting_started_stable_diffusion_with_lora_models/civitai_image_metadata.jpg)

Several articles describe some checkpoint models on the Internet, and there are also [models in a civitai](https://civitai.com/tag/base%20model).
Unlike Lora models, these models are big and take time to download.
I downloaded the next models from Civitai

- [Anything v5](https://civitai.com/models/9409?modelVersionId=30163)
- [YesMix](https://civitai.com/models/9139/checkpointyesmix?modelVersionId=11086)

The checkpoint model should be put under `models/Stable-Diffusion` by following [this step](https://github.com/civitai/civitai/wiki/How-to-use-models#fine-tuned-model-checkpoints-dreambooth-models) and then the model can be selected on the web UI.

### 4.3. Run txt2img

Now is the time to generate images.
At first, I wanted to test checking images and see if the prompts looked good at first.
In order to make it faster,

- Set the width and height lower but sufficient to see the image quality. For my case, set width and height 256.
- Set batch size to 4 to see multiple images at once.

I changed the sampling steps and CFG Scale, but it didn't improve performance, or the quality of images was too bad, and I couldn't change that much.

## 5. Create a new LoRA model

I wanted to test it locally, but it looks like it's easier to create a new model using [a Google Colab notebook](https://colab.research.google.com/github/hollowstrawberry/kohya-colab/blob/main/Lora_Trainer.ipynb#scrollTo=OglZzI_ujZq-).

### 5.1. Prepare a dataset

First, prepare a few images and text files to describe images, and according to [this article](https://anakin.ai/blog/how-to-use-lora-stable-diffusion/), for anime images, booru tags should be used, whereas for other images, the descriptions should be straightforward one.
The example of booru tags can be found [here](https://github.com/DominikDoom/a1111-sd-webui-tagcomplete?tab=readme-ov-file#csv-tag-data), for example.

For the first run, I prepared two image and text files.
Then, set a project name, like `first_lora` in the Google Notebook.
Then, I put the image and text files under the `Loras/first_lora/dataset` folder in my Google Drive.
So, in the end, before running the notebook, the files look like the following:

```bash
> tree /path/to/Google\ Drive/Loras
/path/to/Google Drive/Loras/
└── first_lora
    └── dataset
        ├── 1.jpg
        ├── 1.txt
        ├── 2.jpg
        └── 2.txt
```

### 5.2. Generate a model

On the notebook, click the "Run the cell" button to train a model.
Then multiple models were created under `Loras/first_lora/output`, and the number of models depended on epochs.
Use some of those LoRA models, generate images, and see which one looks good.


## 6. Alternative: Set up a web UI with a conda

There is another way to set up a web UI for WSL2 in [this document](https://github.com/automatic1111/stable-diffusion-webui/wiki/install-and-run-on-nvidia-gpus#windows-11-wsl2-instructions).
This way is basically a manual installation from a source code.
I didn't use this way, though, because it is more complex, and the issues I had were solved using another way to install it.

### 6.1. Install conda

```bash
# install conda (if not already done)
wget https://repo.anaconda.com/archive/Anaconda3-2022.05-Linux-x86_64.sh
chmod +x Anaconda3-2022.05-Linux-x86_64.sh
./Anaconda3-2022.05-Linux-x86_64.sh
$HOME/anaconda3/bin/conda shell.fish hook | source -
```

### 6.2. Set up an environment

```bash
git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui.git
cd stable-diffusion-webui
conda env create -f environment-wsl2.yaml
conda activate automatic
```

### 6.3. Set up a web UI

First, install dependencies for a web UI.
But this might not be necessary because `launch.py` is supposed to install dependencies as well.
And running `python webui.py` wasn't successful.

```bash
mkdir repositories
git clone https://github.com/CompVis/stable-diffusion.git repositories/stable-diffusion-stability-ai
git clone https://github.com/CompVis/taming-transformers.git repositories/taming-transformers
git clone https://github.com/sczhou/CodeFormer.git repositories/CodeFormer
git clone https://github.com/salesforce/BLIP.git repositories/BLIP
pip install transformers==4.19.2 diffusers invisible-watermark --prefer-binary
pip install git+https://github.com/crowsonkb/k-diffusion.git --prefer-binary
pip install git+https://github.com/TencentARC/GFPGAN.git --prefer-binary
pip install -r repositories/CodeFormer/requirements.txt --prefer-binary
pip install -r requirements.txt  --prefer-binary
pip install -U numpy  --prefer-binary
```

Then, launch a web UI using the following command.

```bash
python launch.py
```

## 7. Further readings

- [Hugging Face: Stable Diffusion with Diffusers](https://huggingface.co/blog/stable_diffusion)

## 8. Troubleshooting

### 8.1. Case 1: An error `Cannot locate TCMalloc. Do you have tcmalloc or google-perftool installed on your system? (improves CPU memory usage)`

`google-perftools` needed to be installed.

```bash
> bash webui.sh

################################################################
Install script for stable-diffusion + Web UI
Tested on Debian 11 (Bullseye), Fedora 34+ and openSUSE Leap 15.4 or newer.
################################################################

################################################################
Running on ishikawa user
################################################################

################################################################
Repo already cloned, using it as install directory
################################################################

################################################################
Create and activate python venv
################################################################

################################################################
Launching launch.py...
################################################################
glibc version is 2.35
Cannot locate TCMalloc. Do you have tcmalloc or google-perftool installed on your system? (improves CPU memory usage)
python3: can't open file '/home/ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/launch.py': [Errno 2] No such file or directory
```

### 8.2. Case 2: When there is no `bc` installed (?)

I installed `bc`, and this error was gone. But the `launch.py` was not found anyway.
I moved to a different directory without the `.git` directory, and it started running, though.

```bash
> bash webui.sh

################################################################
Install script for stable-diffusion + Web UI
Tested on Debian 11 (Bullseye), Fedora 34+ and openSUSE Leap 15.4 or newer.
################################################################

################################################################
Running on ishikawa user
################################################################

################################################################
Repo already cloned, using it as install directory
################################################################

################################################################
Create and activate python venv
################################################################

################################################################
Launching launch.py...
################################################################
glibc version is 2.35
Check TCMalloc: libtcmalloc_minimal.so.4
webui.sh: line 251: bc: command not found
webui.sh: line 251: [: -eq: unary operator expected
libtcmalloc_minimal.so.4 is linked with libc.so,execute LD_PRELOAD=/lib/x86_64-linux-gnu/libtcmalloc_minimal.so.4
python3: can't open file '/home/ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/launch.py': [Errno 2] No such file or directory
```

### 8.3. Case 3: RuntimeError: Torch is not able to use GPU; add --skip-torch-cuda-test to COMMANDLINE_ARGS variable to disable this check

After an NVIDIA driver was installed, this error was gone.

```bash
> bash webui.sh

################################################################
Install script for stable-diffusion + Web UI
Tested on Debian 11 (Bullseye), Fedora 34+ and openSUSE Leap 15.4 or newer.
################################################################

################################################################
Running on ishikawa user
################################################################

################################################################
Clone stable-diffusion-webui
################################################################
Cloning into 'stable-diffusion-webui'...
remote: Enumerating objects: 33277, done.
remote: Counting objects: 100% (59/59), done.
remote: Compressing objects: 100% (42/42), done.
remote: Total 33277 (delta 33), reused 32 (delta 14), pack-reused 33218
Receiving objects: 100% (33277/33277), 34.59 MiB | 15.25 MiB/s, done.
Resolving deltas: 100% (23286/23286), done.

################################################################
Create and activate python venv
################################################################

################################################################
Launching launch.py...
################################################################
glibc version is 2.35
Check TCMalloc: libtcmalloc_minimal.so.4
libtcmalloc_minimal.so.4 is linked with libc.so,execute LD_PRELOAD=/lib/x86_64-linux-gnu/libtcmalloc_minimal.so.4
Python 3.10.12 (main, Nov 20 2023, 15:14:05) [GCC 11.4.0]
Version: v1.9.3
Commit hash: 1c0a0c4c26f78c32095ebc7f8af82f5c04fca8c0
Installing torch and torchvision
Looking in indexes: https://pypi.org/simple, https://download.pytorch.org/whl/cu121
Collecting torch==2.1.2
  Downloading https://download.pytorch.org/whl/cu121/torch-2.1.2%2Bcu121-cp310-cp310-linux_x86_64.whl (2200.7 MB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 2.2/2.2 GB 1.5 MB/s eta 0:00:00
Collecting torchvision==0.16.2
  Downloading https://download.pytorch.org/whl/cu121/torchvision-0.16.2%2Bcu121-cp310-cp310-linux_x86_64.whl (6.8 MB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 6.8/6.8 MB 29.4 MB/s eta 0:00:00
Collecting filelock
  Downloading filelock-3.14.0-py3-none-any.whl (12 kB)
Collecting networkx
  Downloading networkx-3.3-py3-none-any.whl (1.7 MB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 1.7/1.7 MB 4.9 MB/s eta 0:00:00
Collecting fsspec
  Downloading fsspec-2024.5.0-py3-none-any.whl (316 kB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 316.1/316.1 KB 9.4 MB/s eta 0:00:00
Collecting typing-extensions
  Downloading typing_extensions-4.12.0-py3-none-any.whl (37 kB)
Collecting triton==2.1.0
  Downloading https://download.pytorch.org/whl/triton-2.1.0-0-cp310-cp310-manylinux2014_x86_64.manylinux_2_17_x86_64.whl (89.2 MB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 89.2/89.2 MB 13.0 MB/s eta 0:00:00
Collecting sympy
  Downloading https://download.pytorch.org/whl/sympy-1.12-py3-none-any.whl (5.7 MB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 5.7/5.7 MB 29.0 MB/s eta 0:00:00
Collecting jinja2
  Downloading jinja2-3.1.4-py3-none-any.whl (133 kB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 133.3/133.3 KB 9.2 MB/s eta 0:00:00
Collecting pillow!=8.3.*,>=5.3.0
  Downloading pillow-10.3.0-cp310-cp310-manylinux_2_28_x86_64.whl (4.5 MB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 4.5/4.5 MB 6.0 MB/s eta 0:00:00
Collecting numpy
  Downloading numpy-1.26.4-cp310-cp310-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (18.2 MB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 18.2/18.2 MB 6.8 MB/s eta 0:00:00
Collecting requests
  Downloading requests-2.32.2-py3-none-any.whl (63 kB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 63.9/63.9 KB 6.3 MB/s eta 0:00:00
Collecting MarkupSafe>=2.0
  Downloading https://download.pytorch.org/whl/MarkupSafe-2.1.5-cp310-cp310-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (25 kB)
Collecting certifi>=2017.4.17
  Downloading certifi-2024.2.2-py3-none-any.whl (163 kB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 163.8/163.8 KB 18.5 MB/s eta 0:00:00
Collecting idna<4,>=2.5
  Downloading idna-3.7-py3-none-any.whl (66 kB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 66.8/66.8 KB 10.2 MB/s eta 0:00:00
Collecting urllib3<3,>=1.21.1
  Downloading urllib3-2.2.1-py3-none-any.whl (121 kB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 121.1/121.1 KB 7.5 MB/s eta 0:00:00
Collecting charset-normalizer<4,>=2
  Downloading charset_normalizer-3.3.2-cp310-cp310-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (142 kB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 142.1/142.1 KB 7.0 MB/s eta 0:00:00
Collecting mpmath>=0.19
  Downloading https://download.pytorch.org/whl/mpmath-1.3.0-py3-none-any.whl (536 kB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 536.2/536.2 KB 31.0 MB/s eta 0:00:00
Installing collected packages: mpmath, urllib3, typing-extensions, sympy, pillow, numpy, networkx, MarkupSafe, idna, fsspec, filelock, charset-normalizer, certifi, triton, requests, jinja2, torch, torchvision
Successfully installed MarkupSafe-2.1.5 certifi-2024.2.2 charset-normalizer-3.3.2 filelock-3.14.0 fsspec-2024.5.0 idna-3.7 jinja2-3.1.4 mpmath-1.3.0 networkx-3.3 numpy-1.26.4 pillow-10.3.0 requests-2.32.2 sympy-1.12 torch-2.1.2+cu121 torchvision-0.16.2+cu121 triton-2.1.0 typing-extensions-4.12.0 urllib3-2.2.1
Traceback (most recent call last):
  File "/home/ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/examples/stable-diffusion/webui/stable-diffusion-webui/launch.py", line 48, in <module>
    main()
  File "/home/ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/examples/stable-diffusion/webui/stable-diffusion-webui/launch.py", line 39, in main
    prepare_environment()
  File "/home/ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/examples/stable-diffusion/webui/stable-diffusion-webui/modules/launch_utils.py", line 386, in prepare_environment
    raise RuntimeError(
RuntimeError: Torch is not able to use GPU; add --skip-torch-cuda-test to COMMANDLINE_ARGS variable to disable this check
```

### 8.4. Case 4: Too slow image generation (Unresolved)

As default, it's suprisingly slow and hard to generate images repeatedly.
I tried to use `--xformers` option to speed it up by following [this comment](https://github.com/AUTOMATIC1111/stable-diffusion-,webui/discussions/5600#discussioncomment-4368337), but it didn't help.
Instead, I just change the resolution to 256x256 to speed it up.

### 8.5. Case 5: Warning: Stable Diffusion XL not found at path /path/to/repositories/generative-models/sgm

When I followed [this document](https://github.com/automatic1111/stable-diffusion-webui/wiki/install-and-run-on-nvidia-gpus#windows-11-wsl2-instructions) to set a web UI up for WSL2, I got an error after running a `webui.py`.
In this case, `launch.py` can be used instead of `webui.py`, which was recommended in [this GitHub comment](https://github.com/AUTOMATIC1111/stable-diffusion-webui/issues/11947#issuecomment-1651791998) to download dependencies.

```bash
> python webui.py                                                                         (automatic)
Warning: Stable Diffusion XL not found at path /home/ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/examples/stable-diffusion/stable-diffusion-webui/repositories/generative-models/sgm
Warning: k_diffusion not found at path /home/ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/examples/stable-diffusion/stable-diffusion-webui/repositories/k-diffusion/k_diffusion/sampling.py
Traceback (most recent call last):
  File "/home/ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/examples/stable-diffusion/stable-diffusion-webui/webui.py", line 13, in <module>
    initialize.imports()
  File "/home/ishikawa/src/github.com/at-ishikawa/at-ishikawa.github.io/examples/stable-diffusion/stable-diffusion-webui/modules/initialize.py", line 32, in imports
    import sgm.modules.encoders.modules  # noqa: F401
ModuleNotFoundError: No module named 'sgm'
```
