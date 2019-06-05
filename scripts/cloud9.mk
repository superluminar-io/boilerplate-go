TMP=/tmp
GOLANG_TARBALL=go1.12.5.linux-amd64.tar.gz

$(TMP)/$(GOLANG_TARBALL):
	@ wget -P $(TMP) https://dl.google.com/go/$(GOLANG_TARBALL)

.PHONY: cloud9
cloud9: $(TMP)/$(GOLANG_TARBALL)
	@ sudo tar -C /usr/local -xzf $<
	@ echo 'export PATH=/usr/local/go/bin:$$HOME/go/bin:$$PATH' >> ~/.bashrc
	@ /usr/local/go/bin/go get github.com/nsf/gocode
	@ sudo cp $$HOME/go/bin/gocode /usr/local/bin
	@ sudo yum install tree -y -q
	@ echo 'Done. Now run the command below.'
	@ echo
	@ echo 'source ~/.bashrc'
	@ echo
	@ echo
