function! whichpr#open_line() abort
  let line = line('.')
  let fname = expand('%')
  let cmd = printf('git blame -L %d,%d %s | cut -d " " -f 1', line, line, fname)
  let sha1 = system(cmd)
  let cmd = printf('whichpr open %s', sha1)
  echo system(cmd)
endfunction

function! whichpr#open_sha1(sha1) abort
  let cmd = printf('whichpr open %s', a:sha1)
  echo system(cmd)
endfunction

function! whichpr#open() abort
  echom "Deprecated. Use whichpr#open_line() instead."
  call whichpr#open_line()
endfunction
