TMP=/tmp
GOLANG_TARBALL=go1.12.5.linux-amd64.tar.gz

$(TMP)/$(GOLANG_TARBALL):
	@ wget -P $(TMP) https://dl.google.com/go/$(GOLANG_TARBALL)

.PHONY: cloud9
cloud9: $(TMP)/$(GOLANG_TARBALL)
	@ sudo tar -C /usr/local -xzf $<
	@ echo 'export PATH=/usr/local/go/bin:$$PATH' >> ~/.bashrc
	@ echo 'Done. Now run the command below.'
	@ echo 'source ~/.bashrc'
