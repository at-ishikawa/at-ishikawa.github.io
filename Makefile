setup:
	bundle install

serve:
	bundle exec jekyll serve --incremental

serve-production:
	JEKYLL_ENV=production make serve
