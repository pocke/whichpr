# whichpr

Find a pull request form commit hash.

## Installation

```bash
go get github.com/pocke/whichpr
```

### OAuth token setting

First time, whichpr asks GitHub username and password.

```bash
$ whichpr open <SHA1>
github.com username: <Enter your username>
github.com password for pocke (never stored): <Enter your password>
two-factor authentication code: <Enter 2FA code if you use 2FA>
```

If you don't want to enter your password, you can register a personal access token manually.

1. Generate a token from [here](https://github.com/settings/tokens/new).
    - `repo` scope is required if you want to access your private repositories.
1. Create `~/.config/whichpr` with the following content.
    ```
    github.com:
    - user: <your username>
      oauth_token: <your token>
      protocol: https
    ```

## Usage

```bash
$ whichpr show SHA1 # => Display a pull request number
$ whichpr open SHA1 # => Open a pull request by your browser
```

### Vim integration

Add the following code to your `.vimrc`.

```vim
set rtp+=~/path/to/pocke/whichpr/

" If you need keybind, please configure yourself.
" For example:
nnoremap <F5> :call whichpr#open()<CR>
```

Execute the following command to open a pull request.

```vim
:call whichpr#open()
```

## Links

- http://qiita.com/pocke/items/281e9157a530a6142178 (Japanese blog)



## License

Copyright 2017-2018 Masataka Kuwabara (pocke)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
