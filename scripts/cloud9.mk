GOLANG_TARBALL = go1.12.5.linux-amd64.tar.gz

/tmp/$(GOLANG_TARBALL):
	wget -P /tmp https://dl.google.com/go/$(GOLANG_TARBALL)

/usr/local/go: /tmp/$(GOLANG_TARBALL)
	sudo tar -C /usr/local -xzf $@ 

cloud9: /usr/local/go
	echo 'PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
	. ~/.bashrc
