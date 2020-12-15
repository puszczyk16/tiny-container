Vagrant.configure("2") do |config|
	config.vm.box = "ubuntu/focal64"
	config.vm.provider "virtualbox" do |vb|
		vb.name = "fbquix_vm"
		vb.memory = "4096"
		vb.cpus = "4"
		# https://github.com/hashicorp/vagrant/issues/11777#issuecomment-661076612
		vb.customize ["modifyvm", :id, "--uart1", "0x3F8", "4"]
		vb.customize ["modifyvm", :id, "--uartmode1", "file", File::NULL]
	end
	config.vm.provision "shell", inline: <<-SHELL
		apt-get update && apt-get upgrade -y
		apt-get install -y curl docker.io snapd git golang python3 python3-pip
		systemctl enable docker
		systemctl start docker
		usermod -aG docker vagrant
		# Install kubectl
		#curl -Lso /tmp/kubectl "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
		#install -m 755 /tmp/kubectl /usr/local/bin
		#curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"	
		SHELL
		config.vm.synced_folder ".", "/syncd"
		config.vm.boot_timeout = 600
end
