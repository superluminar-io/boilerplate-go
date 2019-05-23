GOLANG_TARBALL = go1.12.5.linux-amd64.tar.gz

.PHONY: cloud9
cloud9:
	@ wget -P /tmp https://dl.google.com/go/$(GOLANG_TARBALL)
	@ sudo tar -C /usr/local -xzf /tmp/$(GOLANG_TARBALL)
	@ echo 'export PATH=$$PATH:/usr/local/go/bin' >> ~/.bashrc
	@ echo 'Done. Now run the command below.'
	@ echo 'source ~/.bashrc'
