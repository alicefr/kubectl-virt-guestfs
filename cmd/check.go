package cmd

import (
	"github.com/alicefr/kubectl-virt-guestfs/utils"
	"github.com/spf13/cobra"
	log "k8s.io/klog/v2"
	"os"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if the pvc is in use or not",
	RunE: func(cmd *cobra.Command, args []string) error {
		var inUse bool
		client, err := utils.CreateClient(Config)
		if err != nil {
			return err
		}
		exist, _ := client.ExistsPVC(PvcClaimName, Namespace)
		if !exist {
			log.Infof("The PVC %s doesn't exist", PvcClaimName)
			os.Exit(1)
		}
		inUse, err = client.IsPVCinUse(PvcClaimName, Namespace)
		if err != nil {
			return err
		}

		if inUse {
			log.Infof("The PVC %s is in use", PvcClaimName)
			os.Exit(0)
		}
		log.Infof("PVC %s is not currently used", PvcClaimName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
